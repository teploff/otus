package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
	"github.com/teploff/otus/scheduler/internal/app"
	"github.com/teploff/otus/scheduler/internal/config"
	"github.com/teploff/otus/scheduler/internal/infrastructure/logger"
	"go.uber.org/zap"
)

var (
	configFile = flag.String("config", "./init/config_dev.yaml", "configuration file path")
	dev        = flag.Bool("dev", true, "dev mode")
)

func main() {
	flag.Parse()

	cfg, err := config.LoadFromFile(*configFile)
	if err != nil {
		panic(err)
	}

	zapLogger := logger.New(*dev, &cfg.Logger)
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s pool_max_conns=%d",
		cfg.Db.Username, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name, cfg.Db.SSLMode, cfg.Db.MaxConn)

	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		zapLogger.Fatal("postgres connection error: ", zap.Error(err))
	}
	defer pool.Close()

	stanConn, err := stan.Connect(cfg.Stan.ClusterName, cfg.Stan.ClientID, stan.NatsURL(cfg.Stan.Addr))
	if err != nil {
		zapLogger.Fatal("on stan connect", zap.Error(err))
	}
	zapLogger.Info("connected to stan")
	defer stanConn.Close()

	application := app.NewApp(cfg,
		app.WithLogger(zapLogger),
		app.WithStorage(pool),
		app.WithDataBus(stanConn),
	)

	go application.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	application.Stop()
}
