package main

import (
	"errors"
	"flag"
	"github.com/teploff/otus/calendar/config"
	"github.com/teploff/otus/calendar/endpoint/calendar"
	"github.com/teploff/otus/calendar/infrastructure/logger"
	"github.com/teploff/otus/calendar/internal/implementation/service"
	kitgrpc "github.com/teploff/otus/calendar/transport/grpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"net"
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

	gRPCListener, err := net.Listen("tcp", cfg.GRPCServer.Addr)
	if err != nil {
		zapLogger.Fatal("gRPC listener", zap.Error(err))
	}

	calendarSvc := service.NewCalendarService()

	gRPCServer := kitgrpc.NewGRPCServer(calendar.MakeCalendarEndpoints(calendarSvc),
		logger.NewZapSugarLogger(zapLogger, zapcore.ErrorLevel))

	go func() {
		if err = gRPCServer.Serve(gRPCListener); !errors.Is(err, grpc.ErrServerStopped) && err != nil {
			zapLogger.Fatal("gRPC serve error", zap.Error(err))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	gRPCServer.GracefulStop()
}
