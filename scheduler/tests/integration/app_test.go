//+build integration

package integration

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
	"github.com/stretchr/testify/suite"
	"github.com/teploff/otus/scheduler/internal/app"
	"github.com/teploff/otus/scheduler/internal/config"
)

// Event entity
type Event struct {
	ID               string    `json:"id"`
	ShortDescription string    `json:"short_description"`
	Date             time.Time `json:"date"`
	Duration         int64     `json:"duration"`
	FullDescription  string    `json:"full_description"`
	RemindBefore     int64     `json:"remind_before"`
	UserID           string    `json:"user_id"`
	IsReceived       bool      `json:"is_received"`
}

type ApplicationTestSuite struct {
	suite.Suite
	cfg     config.Config
	app     *app.App
	dbPool  *pgxpool.Pool
	dataBus stan.Conn
}

func (suite *ApplicationTestSuite) SetupSuite() {
	dbPort, err := strconv.Atoi(os.Getenv("TEST_DB_PORT"))
	suite.NoError(err)

	cfg := config.Config{
		Db: config.DbConfig{
			Host:     os.Getenv("TEST_DB_HOST"),
			Port:     dbPort,
			Name:     os.Getenv("TEST_DB_NAME"),
			Username: os.Getenv("TEST_DB_USER"),
			Password: os.Getenv("TEST_DB_PASSWORD"),
			SSLMode:  os.Getenv("TEST_DB_SSL_MODE"),
			MaxConn:  10,
		},
		Stan: config.StanConfig{
			ClusterName: os.Getenv("TEST_STAN_CLUSTER_NAME"),
			ClientID:    os.Getenv("TEST_STAN_CLIENT_ID"),
			Addr:        os.Getenv("TEST_STAN_ADDR"),
		},
		Scheduler: config.SchedulerConfig{Interval: time.Millisecond * 500},
	}
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s pool_max_conns=%d",
		cfg.Db.Username, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name, cfg.Db.SSLMode, cfg.Db.MaxConn)

	pool, err := pgxpool.Connect(context.Background(), dsn)
	suite.NoError(err)

	stanConn, err := stan.Connect(cfg.Stan.ClusterName, cfg.Stan.ClientID, stan.NatsURL(cfg.Stan.Addr))
	suite.NoError(err)

	application := app.NewApp(cfg,
		app.WithStorage(pool),
		app.WithDataBus(stanConn),
	)

	go application.Run()

	suite.cfg = cfg
	suite.app = application
	suite.dbPool = pool
	suite.dataBus = stanConn
}

func (suite *ApplicationTestSuite) TearDownTest() {
	ctx := context.Background()
	tx, _ := suite.dbPool.Begin(ctx)
	defer tx.Rollback(ctx)

	_, err := tx.Exec(ctx, `
	TRUNCATE TABLE public."Event"`)
	suite.NoError(err)

	suite.NoError(tx.Commit(ctx))
}

func (suite *ApplicationTestSuite) TearDownSuite() {
	suite.app.Stop()
	suite.dbPool.Close()
	suite.NoError(suite.dataBus.Close())
}

// TestFindNotificationsAndSendViaStan checks that application finds notifications which must be send via data-bus.
func (suite *ApplicationTestSuite) TestFindNotificationsAndSendViaStan() {
	suite.saveEvents()

	counter := suite.subscribeEvents()
	suite.Equal(int32(2), counter)

	events := suite.getEvents()
	suite.Equal(3, len(events))
	suite.True(events[0].IsReceived)
	suite.True(events[1].IsReceived)
	suite.False(events[2].IsReceived)
}

func (suite *ApplicationTestSuite) saveEvents() {
	ctx := context.Background()
	tx, err := suite.dbPool.Begin(ctx)
	suite.NoError(err)
	defer tx.Rollback(ctx)

	_, err = tx.Exec(context.Background(), `
	INSERT
		INTO public."Event"(short_description, date, duration, full_description, remind_before, user_id)
	VALUES
		($1, $2, $3, $4, $5, $6)`, "short_desc_1",
		time.Now().UTC().Add(time.Minute*2).Truncate(time.Second).Truncate(time.Minute), 12, "full_desc_1", 1,
		uuid.NewV4())
	suite.NoError(err)

	_, err = tx.Exec(context.Background(), `
	INSERT
		INTO public."Event"(short_description, date, duration, full_description, remind_before, user_id)
	VALUES
		($1, $2, $3, $4, $5, $6)`, "short_desc_2",
		time.Now().UTC().Add(time.Minute*2).Truncate(time.Second).Truncate(time.Minute), 12, "full_desc_2", 1,
		uuid.NewV4())
	suite.NoError(err)

	_, err = tx.Exec(context.Background(), `
	INSERT
		INTO public."Event"(short_description, date, duration, full_description, remind_before, user_id)
	VALUES
		($1, $2, $3, $4, $5, $6)`, "short_desc_3", time.Now().UTC().Add(time.Minute*5), 12, "full_desc_3",
		1, uuid.NewV4())
	suite.NoError(err)

	suite.NoError(tx.Commit(ctx))
}

func (suite *ApplicationTestSuite) subscribeEvents() int32 {
	var counter int32
	sc1, err := stan.Connect(suite.cfg.Stan.ClusterName, uuid.NewV4().String(),
		stan.NatsURL(suite.cfg.Stan.Addr))
	suite.NoError(err)

	var sub stan.Subscription
	go func() {
		sub, err = sc1.Subscribe("notifications",
			func(m *stan.Msg) {
				suite.NoError(m.Ack())
				atomic.AddInt32(&counter, 1)
			},
			stan.SetManualAckMode(),
		)
		suite.NoError(err)
		suite.NoError(sub.SetPendingLimits(-1, -1))
	}()
	<-time.After(time.Minute * 2)

	suite.NoError(sc1.Close())
	return counter
}

func (suite *ApplicationTestSuite) getEvents() []Event {
	ctx := context.Background()
	tx, err := suite.dbPool.Begin(ctx)
	suite.NoError(err)
	defer tx.Rollback(ctx)

	events := make([]Event, 0, 100)
	rows, err := tx.Query(ctx, `
	SELECT id, short_description, date, duration, full_description, remind_before, user_id, is_received
	FROM public."Event"
	ORDER BY date ASC`)
	suite.NoError(err)
	defer rows.Close()

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.ShortDescription, &event.Date, &event.Duration, &event.FullDescription,
			&event.RemindBefore, &event.UserID, &event.IsReceived)
		suite.NoError(err)
		events = append(events, event)
	}

	suite.NoError(tx.Commit(ctx))

	return events
}

func TestApplication(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuite))
}
