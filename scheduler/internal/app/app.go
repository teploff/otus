package app

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
	"github.com/teploff/otus/scheduler/internal/config"
	"github.com/teploff/otus/scheduler/internal/implementation/databus"
	"github.com/teploff/otus/scheduler/internal/implementation/repository"
	"github.com/teploff/otus/scheduler/internal/implementation/service"
	"go.uber.org/zap"
)

// AppOption via application.
type Option func(*App)

// WithLogger adding logger option.
func WithLogger(l *zap.Logger) Option {
	return func(a *App) {
		a.logger = l
	}
}

// WithDataBus adding DataBus (nats-streaming) option.
func WithDataBus(conn stan.Conn) Option {
	return func(a *App) {
		a.stanConn = conn
	}
}

// WithStorage adding storage (Postgres) option.
func WithStorage(pool *pgxpool.Pool) Option {
	return func(a *App) {
		a.storagePool = pool
	}
}

type App struct {
	cfg             config.Config
	stanConn        stan.Conn
	storagePool     *pgxpool.Pool
	logger          *zap.Logger
	stopCommandChan chan struct{}
}

// NewApp returns instance of app.
func NewApp(cfg config.Config, opts ...Option) *App {
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

// Run lunch application.
func (a *App) Run() {
	storage := repository.NewEventRepository(a.storagePool)
	stanService := databus.NewStanDataBus(a.stanConn)
	reminderService := service.NewTickerReminder(a.cfg.Scheduler.Interval, storage, stanService, a.logger)

	ctx, cancel := context.WithCancel(context.Background())
	go reminderService.Run(ctx)

	<-a.stopCommandChan

	cancel()
	reminderService.Stop()
}

// Stop gracefully shutting down application.
func (a *App) Stop() {
	a.stopCommandChan <- struct{}{}
}
