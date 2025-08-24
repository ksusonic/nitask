package ticket

import (
	"context"
	"log/slog"

	"github.com/ksusonic/nitask/internal/models"
)

func (c *Controller) Create(ctx context.Context, in models.TicketCreateIn) (*models.Ticket, error) {
	ticket, err := c.repository.Create(ctx, in)
	if err != nil {
		c.log.ErrorContext(ctx, "create ticket", slog.Any("error", err))

		return nil, err
	}

	return ticket, nil
}

func (c *Controller) Get(ctx context.Context, key string) (*models.Ticket, error) {
	ticket, err := c.repository.Get(ctx, key)
	if err != nil {
		c.log.ErrorContext(ctx, "get ticket", slog.Any("error", err))

		return nil, err
	}

	return ticket, nil
}

func (c *Controller) List(ctx context.Context, in models.TicketListIn) ([]models.Ticket, error) {
	tickets, err := c.repository.List(ctx, in)
	if err != nil {
		c.log.ErrorContext(ctx, "list ticket", slog.Any("error", err))

		return nil, err
	}

	return tickets, nil
}

func (c *Controller) Update(ctx context.Context, key string, in models.TicketUpdateIn) (*models.Ticket, error) {
	ticket, err := c.repository.Update(ctx, key, in)
	if err != nil {
		c.log.ErrorContext(ctx, "update ticket", slog.Any("error", err))

		return nil, err
	}

	return ticket, nil
}

func (c *Controller) Delete(ctx context.Context, key string) error {
	err := c.repository.Delete(ctx, key)
	if err != nil {
		c.log.ErrorContext(ctx, "delete ticket", slog.Any("error", err))

		return err
	}

	return nil
}
