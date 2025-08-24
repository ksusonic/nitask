package app

import (
	"context"
	"fmt"
	"log/slog"

	ticketController "github.com/ksusonic/nitask/internal/controller/ticket"
	ticketRepo "github.com/ksusonic/nitask/internal/repository/ticket"
	"github.com/ksusonic/nitask/internal/storage"
	"github.com/ksusonic/nitask/pkg/config"
	"github.com/ksusonic/nitask/pkg/logger"
)

type App struct {
	config *config.Config
	log    *slog.Logger
	mongo  *storage.Mongo

	ticketController *ticketController.Controller
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	log := logger.NewLogger(cfg.Logger)

	mongo, err := storage.NewMongo(cfg.MongoDB)
	if err != nil {
		return nil, fmt.Errorf("init storage: %w", err)
	}

	ticketStorage, err := ticketRepo.NewRepository(ctx, mongo.Client(), true)
	if err != nil {
		return nil, fmt.Errorf("init ticket storage: %w", err)
	}

	return &App{
		config:           cfg,
		log:              log,
		mongo:            mongo,
		ticketController: ticketController.New(ticketStorage, log),
	}, nil
}

func (a *App) Config() *config.Config {
	return a.config
}

func (a *App) Logger() *slog.Logger {
	return a.log
}

func (a *App) MongoDB() *storage.Mongo {
	return a.mongo
}

func (a *App) TicketController() *ticketController.Controller {
	return a.ticketController
}

func (a *App) Close(ctx context.Context) error {
	if err := a.mongo.Close(ctx); err != nil {
		return fmt.Errorf("close mongo: %w", err)
	}

	a.log.DebugContext(ctx, "closed mongo")

	return nil
}
