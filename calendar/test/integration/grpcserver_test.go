package integration

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"github.com/teploff/otus/calendar/config"
	"github.com/teploff/otus/calendar/pkg"
	"github.com/teploff/otus/calendar/transport/grpc/pb"
	"google.golang.org/grpc"
	"testing"
)

type GRPCServerTS struct {
	suite.Suite
	connPool *pgxpool.Pool
	addr     string
	app      *pkg.App
}

func (g *GRPCServerTS) SetupSuite() {
	cfg, err := config.LoadFromFile("../../init/config_test.yaml")
	if err != nil {
		panic(err)
	}
	g.addr = cfg.GRPCServer.Addr

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s pool_max_conns=%d",
		cfg.Db.Username, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name, cfg.Db.SSLMode, cfg.Db.MaxConn)
	g.connPool, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	g.app = pkg.NewApp(cfg,
		pkg.WithConnPool(g.connPool),
	)

	go g.app.Run()
}

func (g *GRPCServerTS) TearDownSuite() {
	defer g.connPool.Close()
}

func (g *GRPCServerTS) TearDownTest() {
	g.flushDB()
}

func TestGRPCServer(t *testing.T) {
	suite.Run(t, new(GRPCServerTS))
}

func (g *GRPCServerTS) TestUpdateEvent() {
	conn, err := grpc.Dial(g.addr, grpc.WithInsecure())
	g.Assert().NoError(err)

	client := pb.NewCalendarClient(conn)

	userID := uuid.NewV4()
	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 1",
			// Sun, 10 May 2020 01:00:00 GMT
			Date:            1589072400,
			Duration:        40,
			FullDescription: "Full description of event 1",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	resp, err := client.GetDailyEvent(context.Background(), &pb.DateRequest{
		UserId: userID.String(),
		// Sun, 10 May 2020 00:00:00 GMT
		Date: 1589068800,
	})
	g.Assert().NoError(err)
	g.Assert().Equal(1, len(resp.Events))
	eventID := resp.Events[0].Id

	_, err = client.UpdateEvent(context.Background(), &pb.UpdateEventRequest{
		UserId:  userID.String(),
		EventId: eventID,
		Event: &pb.Event{
			ShortDescription: "new event 1",
			Date:             1589072400,
			Duration:         50,
			FullDescription:  "New Full description of event 1",
			RemindBefore:     20,
		},
	})
	g.Assert().NoError(err)
	resp, err = client.GetDailyEvent(context.Background(), &pb.DateRequest{
		UserId: userID.String(),
		// Sun, 10 May 2020 00:00:00 GMT
		Date: 1589068800,
	})
	g.Assert().NoError(err)
	g.Assert().Equal(1, len(resp.Events))
	g.Assert().Equal(resp.Events[0].ShortDescription, "new event 1")
	g.Assert().Equal(resp.Events[0].Date, int64(1589072400))
	g.Assert().Equal(resp.Events[0].FullDescription, "New Full description of event 1")
	g.Assert().Equal(resp.Events[0].Duration, int64(50))
	g.Assert().Equal(resp.Events[0].RemindBefore, int64(20))
}

func (g *GRPCServerTS) TestDeleteEvent() {
	conn, err := grpc.Dial(g.addr, grpc.WithInsecure())
	g.Assert().NoError(err)

	client := pb.NewCalendarClient(conn)

	userID := uuid.NewV4()
	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 1",
			// Sun, 10 May 2020 01:00:00 GMT
			Date:            1589072400,
			Duration:        40,
			FullDescription: "Full description of event 1",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	resp, err := client.GetDailyEvent(context.Background(), &pb.DateRequest{
		UserId: userID.String(),
		// Sun, 10 May 2020 00:00:00 GMT
		Date: 1589068800,
	})
	g.Assert().NoError(err)
	g.Assert().Equal(1, len(resp.Events))
	eventID := resp.Events[0].Id

	_, err = client.DeleteEvent(context.Background(), &pb.DeleteEventRequest{
		UserId:  userID.String(),
		EventId: eventID,
	})
	g.Assert().NoError(err)
	resp, err = client.GetDailyEvent(context.Background(), &pb.DateRequest{
		UserId: userID.String(),
		// Sun, 10 May 2020 00:00:00 GMT
		Date: 1589068800,
	})
	g.Assert().NoError(err)
	g.Assert().Equal(0, len(resp.Events))
}

func (g *GRPCServerTS) TestGetDailyEvent() {
	conn, err := grpc.Dial(g.addr, grpc.WithInsecure())
	g.Assert().NoError(err)

	client := pb.NewCalendarClient(conn)

	userID := uuid.NewV4()
	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 1",
			// Sun, 10 May 2020 01:00:00 GMT
			Date:            1589072400,
			Duration:        40,
			FullDescription: "Full description of event 1",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 2",
			// Sun, 10 May 2020 22:00:00 GMT
			Date:            1589148000,
			Duration:        40,
			FullDescription: "Full description of event 2",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 3",
			// Sun, 11 May 2020 01:00:00 GMT
			Date:            1589158800,
			Duration:        40,
			FullDescription: "Full description of event 3",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	resp, err := client.GetDailyEvent(context.Background(), &pb.DateRequest{
		UserId: userID.String(),
		// Sun, 10 May 2020 00:00:00 GMT
		Date: 1589068800,
	})
	g.Assert().NoError(err)
	g.Assert().Equal(2, len(resp.Events))
	g.Assert().Equal("event 1", resp.Events[0].ShortDescription)
	g.Assert().Equal("event 2", resp.Events[1].ShortDescription)
	g.Assert().Equal("Full description of event 1", resp.Events[0].FullDescription)
	g.Assert().Equal("Full description of event 2", resp.Events[1].FullDescription)
}

func (g *GRPCServerTS) TestGetWeeklyEvent() {
	conn, err := grpc.Dial(g.addr, grpc.WithInsecure())
	g.Assert().NoError(err)

	client := pb.NewCalendarClient(conn)

	userID := uuid.NewV4()
	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 1",
			// Sun, 11 May 2020 01:00:00 GMT
			Date:            1589158800,
			Duration:        40,
			FullDescription: "Full description of event 1",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 2",
			// Sun, 17 May 2020 01:00:00 GMT
			Date:            1589677200,
			Duration:        40,
			FullDescription: "Full description of event 2",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 3",
			// Sun, 18 May 2020 10:00:00 GMT
			Date:            1589796000,
			Duration:        40,
			FullDescription: "Full description of event 3",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	resp, err := client.GetWeeklyEvent(context.Background(), &pb.DateRequest{
		UserId: userID.String(),
		// Sun, 11 May 2020 00:00:00 GMT
		Date: 1589155200,
	})
	g.Assert().NoError(err)
	g.Assert().Equal(2, len(resp.Events))
	g.Assert().Equal("event 1", resp.Events[0].ShortDescription)
	g.Assert().Equal("event 2", resp.Events[1].ShortDescription)
	g.Assert().Equal("Full description of event 1", resp.Events[0].FullDescription)
	g.Assert().Equal("Full description of event 2", resp.Events[1].FullDescription)
}

func (g *GRPCServerTS) TestGetMonthlyEvent() {
	conn, err := grpc.Dial(g.addr, grpc.WithInsecure())
	g.Assert().NoError(err)

	client := pb.NewCalendarClient(conn)

	userID := uuid.NewV4()
	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 1",
			// Sun, 15 May 2020 10:00:00 GMT
			Date:            1589536800,
			Duration:        40,
			FullDescription: "Full description of event 1",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 2",
			// Sun, 29 May 2020 16:30:00 GMT
			Date:            1590769800,
			Duration:        40,
			FullDescription: "Full description of event 2",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	_, err = client.CreateEvent(context.Background(), &pb.CreateRequest{
		UserId: userID.String(),
		Event: &pb.Event{
			ShortDescription: "event 3",
			// Sun, 29 June 2020 16:30:00 GMT
			Date:            1593448200,
			Duration:        40,
			FullDescription: "Full description of event 3",
			RemindBefore:    30,
		},
	})
	g.Assert().NoError(err)

	resp, err := client.GetMonthlyEvent(context.Background(), &pb.DateRequest{
		UserId: userID.String(),
		// Sun, 1 May 2020 00:00:00 GMT
		Date: 1588291200,
	})
	g.Assert().NoError(err)
	g.Assert().Equal(2, len(resp.Events))
	g.Assert().Equal("event 1", resp.Events[0].ShortDescription)
	g.Assert().Equal("event 2", resp.Events[1].ShortDescription)
	g.Assert().Equal("Full description of event 1", resp.Events[0].FullDescription)
	g.Assert().Equal("Full description of event 2", resp.Events[1].FullDescription)
}

func (g *GRPCServerTS) flushDB() {
	ctx := context.Background()
	tx, _ := g.connPool.Begin(ctx)
	defer tx.Rollback(ctx)

	_, _ = tx.Exec(ctx, `
	TRUNCATE TABLE public."Event"`)

	_ = tx.Commit(ctx)
}
