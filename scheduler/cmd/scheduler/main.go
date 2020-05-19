package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
	"github.com/teploff/otus/scheduler/internal/config"
	"github.com/teploff/otus/scheduler/internal/implementation/databus"
	"github.com/teploff/otus/scheduler/internal/implementation/repository"
	"github.com/teploff/otus/scheduler/internal/implementation/service"
	"github.com/teploff/otus/scheduler/internal/infrastructure/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
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
	storage := repository.NewEventRepository(pool)

	stanConn, err := stan.Connect(cfg.Stan.ClusterName, cfg.Stan.ClientID, stan.NatsURL(cfg.Stan.Addr))
	if err != nil {
		zapLogger.Fatal("on stan connect", zap.Error(err))
	}
	zapLogger.Info("connected to stan")
	defer stanConn.Close()

	stanService := databus.NewStanDataBus(stanConn)

	reminderService := service.NewTickerReminder(cfg.Scheduler.Interval, storage, stanService, zapLogger)
	ctx, cancel := context.WithCancel(context.Background())
	reminderService.Run(ctx)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	cancel()
	reminderService.Stop()
}
