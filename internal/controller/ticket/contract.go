package ticket

import (
	"context"

	"github.com/ksusonic/nitask/internal/models"
)

type Repository interface {
	List(ctx context.Context, in models.TicketListIn) ([]models.Ticket, error)
	Get(ctx context.Context, key string) (*models.Ticket, error)
	Create(ctx context.Context, in models.TicketCreateIn) (*models.Ticket, error)
	Update(ctx context.Context, key string, in models.TicketUpdateIn) (*models.Ticket, error)
	Delete(ctx context.Context, key string) error
}
