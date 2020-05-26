//+build integration

package integration

import (
	"github.com/nats-io/stan.go"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/teploff/otus/sender/internal/app"
	"github.com/teploff/otus/sender/internal/config"
	"github.com/teploff/otus/sender/internal/enpoints/notifier"
	"os"
	"strconv"
	"testing"
	"time"
)

type loggerMock struct {
	mock.Mock
}

func (l *loggerMock) Log(keyvals ...interface{}) error {
	l.Called(keyvals)
	return nil
}

func createLoggerMock() *loggerMock {
	lg := &loggerMock{}
	lg.On("Log", mock.Anything).Return(nil)
	return lg
}

var mockLogger = createLoggerMock()

type ApplicationTestSuite struct {
	suite.Suite
	cfg     config.Config
	app     *app.App
	l       mock.Mock
	dataBus stan.Conn
}

func (suite *ApplicationTestSuite) SetupSuite() {
	cfg := config.Config{
		Stan: config.StanConfig{
			ClusterName: os.Getenv("TEST_STAN_CLUSTER_NAME"),
			ClientID:    os.Getenv("TEST_STAN_CLIENT_ID"),
			Addr:        os.Getenv("TEST_STAN_ADDR"),
		},
	}

	stanConn, err := stan.Connect(cfg.Stan.ClusterName, cfg.Stan.ClientID, stan.NatsURL(cfg.Stan.Addr))
	suite.NoError(err)

	application := app.NewApp(cfg,
		app.WithLogger(mockLogger),
		app.WithDataBus(stanConn),
	)

	go application.Run()

	suite.cfg = cfg
	suite.app = application
	suite.dataBus = stanConn
}

func (suite *ApplicationTestSuite) TearDownSuite() {
	suite.app.Stop()
	suite.NoError(suite.dataBus.Close())
}

// TestGetNotificationsViaStan checks that application getting notifications via data-bus and log in zap.
func (suite *ApplicationTestSuite) TestGetNotificationsViaStan() {
	suite.publishSysMetrics()
	<-time.After(time.Second * 2)
	mockLogger.AssertNumberOfCalls(suite.T(), "Log", 20)
}

func (suite *ApplicationTestSuite) publishSysMetrics() {
	sc, err := stan.Connect(suite.cfg.Stan.ClusterName, uuid.NewV4().String(), stan.NatsURL(suite.cfg.Stan.Addr),
		stan.MaxPubAcksInflight(1000))
	suite.NoError(err)

	for i := 0; i < 20; i++ {
		request := notifier.SendNotificationRequest{
			UserID:           uuid.NewV4().String(),
			ShortDescription: "short_description_" + strconv.Itoa(i),
			FullDescription:  "full_description_" + strconv.Itoa(i),
			Date:             time.Now(),
			Duration:         int64(i),
		}
		msg, err := request.MarshalJSON()
		suite.NoError(err)
		_, err = sc.PublishAsync("notifications", msg, func(_ string, _ error) {})
		suite.NoError(err)
	}
	suite.NoError(sc.Close())
}

func TestApplication(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuite))
}
