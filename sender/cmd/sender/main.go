package main

import (
	"flag"
	"github.com/nats-io/stan.go"
	"github.com/teploff/otus/sender/internal/config"
	"github.com/teploff/otus/sender/internal/enpoints/notifier"
	"github.com/teploff/otus/sender/internal/implementation/service"
	"github.com/teploff/otus/sender/internal/infrastructure/logger"
	kitstan "github.com/teploff/otus/sender/internal/transport/stan"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	svc := service.NewNotifierService(zapLogger)

	stanServer := kitstan.NewStan()
	go func() {
		err = stanServer.Serve(stanConn, notifier.MakeNotifierEndpoints(svc), logger.NewZapSugarLogger(zapLogger, zapcore.ErrorLevel))
		if err != nil {
			zapLogger.Fatal("stan serve error", zap.Error(err))
		}
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	stanServer.Stop()
}
