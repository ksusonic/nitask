package handler

import (
	"context"
	"log/slog"

	"github.com/ksusonic/nitask/internal/models"
)

type Deps struct {
	TicketController ticketController
	Logger           *slog.Logger
}

type ticketController interface {
	Create(ctx context.Context, in models.TicketCreateIn) (*models.Ticket, error)
	Get(ctx context.Context, key string) (*models.Ticket, error)
	List(ctx context.Context, in models.TicketListIn) ([]models.Ticket, error)
	Update(ctx context.Context, key string, in models.TicketUpdateIn) (*models.Ticket, error)
	Delete(ctx context.Context, key string) error
}
