package main

import (
	"context"
	"flag"
	"github.com/otus/calendar/configs"
	"github.com/otus/calendar/infrastrucure/logger"
	"github.com/otus/calendar/infrastrucure/server"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	pathConfiguration = flag.String("config", "./init/config_dev.yaml", "config file path")
)

func main() {
	flag.Parse()

	cfg, err := configs.LoadConfiguration(*pathConfiguration)
	if err != nil {
		panic(err)
	}

	logger.InitZapLogger(cfg.Logger)

	srv := server.NewHTTPServer(cfg.Server)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal(err.Error())
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err = srv.Shutdown(ctx); err != nil {
		zap.L().Fatal(err.Error())
	}

	zap.L().Info("Http server is gracefully shutdown")
}
