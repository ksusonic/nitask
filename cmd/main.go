package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/ksusonic/nitask/internal/app"
	"github.com/ksusonic/nitask/pkg/config"
	"github.com/ksusonic/nitask/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	log := logger.NewLogger(cfg.Logger)

	application, err := app.New(cfg, log)
	if err != nil {
		log.Error("init app", slog.Any("error", err))
		os.Exit(1)
	}

	application.Run(context.Background())
}
