package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ksusonic/nitask/internal/handler"
	"github.com/ksusonic/nitask/internal/server"
	"github.com/ksusonic/nitask/internal/storage"
	"github.com/ksusonic/nitask/pkg/config"
)

type App struct {
	config *config.Config
	server *server.Server
	mongo  *storage.Mongo
	log    *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) (*App, error) {
	srv := server.New(cfg.Server)

	mongo, err := storage.NewMongo(cfg.MongoDB)
	if err != nil {
		return nil, fmt.Errorf("init storage: %w", err)
	}

	handler.RegisterRoutes(srv.Engine)

	return &App{
		config: cfg,
		server: srv,
		mongo:  mongo,
		log:    log,
	}, nil
}

func (a *App) Run(ctx context.Context) {
	server := &http.Server{
		Addr:    a.config.Server.Address,
		Handler: a.server,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			a.log.Error("server closing", slog.Any("error", err))
		}
	}()

	<-quit
	a.log.Info("interrupt signal")

	ctx, timeout := context.WithTimeout(ctx, 3*time.Second)
	defer timeout()

	shutdown := make(chan struct{})
	go func() {
		defer close(shutdown)

		if err := server.Close(); err != nil {
			a.log.Error("server close", slog.Any("error", err))
		}

		if err := a.Close(ctx); err != nil {
			a.log.Error("close app", slog.Any("error", err))
		}
	}()

	select {
	case <-shutdown:
		a.log.Debug("app gracefully closed")
	case <-ctx.Done():
		a.log.Error("shutdown timeout")
	}
}

func (a *App) Close(ctx context.Context) error {
	if err := a.mongo.Close(ctx); err != nil {
		return fmt.Errorf("close mongo: %w", err)
	}

	a.log.Debug("closed mongo")

	return nil
}
