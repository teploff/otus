package pkg

import (
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/teploff/otus/calendar/config"
	"github.com/teploff/otus/calendar/endpoint/calendar"
	"github.com/teploff/otus/calendar/infrastructure/logger"
	"github.com/teploff/otus/calendar/internal/implementation/repository"
	"github.com/teploff/otus/calendar/internal/implementation/service"
	kitgrpc "github.com/teploff/otus/calendar/transport/grpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"net"
)

// AppOption via application
type AppOption func(*App)

// WithLogger adding logger option
func WithLogger(l *zap.Logger) AppOption {
	return func(a *App) {
		a.logger = l
	}
}

// WithConnPool adding psql pool connections
func WithConnPool(pool *pgxpool.Pool) AppOption {
	return func(a *App) {
		a.poolConn = pool
	}
}

// App is application to encapsulate login to launch in main
type App struct {
	cfg             config.Config
	poolConn        *pgxpool.Pool
	logger          *zap.Logger
	stopCommandChan chan struct{}
}

// NewApp returns instance of app
func NewApp(cfg config.Config, opts ...AppOption) *App {
	app := &App{
		cfg:             cfg,
		logger:          zap.NewNop(),
		stopCommandChan: make(chan struct{}),
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

// Run lunch application
func (a *App) Run() {
	gRPCListener, err := net.Listen("tcp", a.cfg.GRPCServer.Addr)
	if err != nil {
		a.logger.Fatal("gRPC transport", zap.Error(err))
	}

	calendarSvc := service.NewCalendarService(repository.NewEventRepository(a.poolConn))

	gRPCServer := kitgrpc.NewGRPCServer(calendar.MakeCalendarEndpoints(calendarSvc),
		logger.NewZapSugarLogger(a.logger, zapcore.ErrorLevel))

	go func() {
		if err = gRPCServer.Serve(gRPCListener); !errors.Is(err, grpc.ErrServerStopped) && err != nil {
			a.logger.Fatal("gRPC serve error", zap.Error(err))
		}
	}()

	<-a.stopCommandChan
	gRPCServer.GracefulStop()
}

// Stop gracefully shutting down application
func (a *App) Stop() {
	a.stopCommandChan <- struct{}{}
}
