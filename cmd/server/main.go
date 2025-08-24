package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ksusonic/nitask/internal/app"
	"github.com/ksusonic/nitask/internal/server"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatalf("init app: %v", err)
	}

	os.Exit(run(app))
}

func run(app *app.App) int {
	ctx := context.Background()

	server := &http.Server{
		Addr:    app.Config().Server.Address,
		Handler: server.New(app.Handler()),
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			app.Logger().Error("server closing", slog.Any("error", err))
		}
	}()

	<-quit
	app.Logger().Info("interrupt signal")

	ctx, timeout := context.WithTimeout(ctx, 3*time.Second)
	defer timeout()

	shutdown := make(chan struct{})
	go func() {
		defer close(shutdown)

		if err := server.Close(); err != nil {
			app.Logger().Error("server close", slog.Any("error", err))
		}

		if err := app.Close(ctx); err != nil {
			app.Logger().Error("close app", slog.Any("error", err))
		}
	}()

	select {
	case <-shutdown:
		app.Logger().Debug("app gracefully closed")
	case <-ctx.Done():
		app.Logger().Error("shutdown timeout")
	}

	return 0
}
