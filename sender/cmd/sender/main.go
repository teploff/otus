package main

import (
	"flag"
	"github.com/nats-io/stan.go"
	"github.com/teploff/otus/sender/internal/app"
	"github.com/teploff/otus/sender/internal/config"
	"github.com/teploff/otus/sender/internal/infrastructure/logger"
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

	stanConn, err := stan.Connect(cfg.Stan.ClusterName, cfg.Stan.ClientID, stan.NatsURL(cfg.Stan.Addr))
	if err != nil {
		zapLogger.Fatal("stan transport", zap.Error(err))
	}
	defer stanConn.Close()

	application := app.NewApp(cfg,
		app.WithLogger(zapLogger),
		app.WithDataBus(stanConn),
	)

	go application.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	application.Stop()
}
