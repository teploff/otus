package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/teploff/otus/calendar/config"
	"github.com/teploff/otus/calendar/infrastructure/logger"
	"github.com/teploff/otus/calendar/pkg"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var (
	configFile = flag.String("config", "./init/config_dev.yaml", "configuration file path")
	dev        = flag.Bool("dev", false, "dev mode")
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

	app := pkg.NewApp(cfg,
		pkg.WithLogger(zapLogger),
		pkg.WithConnPool(pool),
	)

	go app.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	app.Stop()
}
