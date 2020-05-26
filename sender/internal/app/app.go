package app

import (
	"github.com/go-kit/kit/log"
	"github.com/nats-io/stan.go"
	"github.com/teploff/otus/sender/internal/config"
	"github.com/teploff/otus/sender/internal/enpoints/notifier"
	"github.com/teploff/otus/sender/internal/implementation/service"
	kitstan "github.com/teploff/otus/sender/internal/transport/stan"
	"go.uber.org/zap"
)

// AppOption via application.
type Option func(*App)

// WithLogger adding logger option.
func WithLogger(l log.Logger) Option {
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

type App struct {
	cfg             config.Config
	stanConn        stan.Conn
	logger          log.Logger
	stopCommandChan chan struct{}
}

// NewApp returns instance of app.
func NewApp(cfg config.Config, opts ...Option) *App {
	app := &App{
		cfg:             cfg,
		stopCommandChan: make(chan struct{}),
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

// Run lunch application.
func (a *App) Run() {
	svc := service.NewNotifierService(a.logger)

	stanServer := kitstan.NewStan()
	go func() {
		err := stanServer.Serve(a.stanConn, notifier.MakeNotifierEndpoints(svc), a.logger)
		if err != nil {
			a.logger.Log("stan serve error", zap.Error(err))
		}
	}()

	<-a.stopCommandChan

	stanServer.Stop()
}

// Stop gracefully shutting down application.
func (a *App) Stop() {
	a.stopCommandChan <- struct{}{}
}
