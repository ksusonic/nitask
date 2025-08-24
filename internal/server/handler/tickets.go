package handler

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"github.com/ksusonic/nitask/internal/models"
	"github.com/ksusonic/nitask/pkg/api"
)

func (h *Handler) GetTickets(
	ctx context.Context,
	req api.GetTicketsRequestObject,
) (api.GetTicketsResponseObject, error) {
	limit := int64(pointer.Get(req.Params.Limit))
	if limit == 0 {
		limit = 10
	}

	tickets, err := h.TicketController.List(ctx, models.TicketListIn{
		Queue:  req.Params.Queue,
		Offset: int64(pointer.Get(req.Params.Offset)),
		Limit:  limit,
	})
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}

	response := make(api.GetTickets200JSONResponse, 0, len(tickets))
	for _, ticket := range tickets {
		response = append(response, api.Ticket{
			CreatedAt:   ticket.CreatedAt,
			Description: ticket.Description,
			Key:         ticket.Key,
			Status:      api.TicketStatus(ticket.Status),
			Title:       ticket.Title,
			UpdatedAt:   ticket.UpdatedAt,
		})
	}

	return response, nil
}

func (h *Handler) PostTickets(
	ctx context.Context,
	req api.PostTicketsRequestObject,
) (api.PostTicketsResponseObject, error) {
	var idempotencyKey uuid.UUID

	if req.Body.IdempotencyKey == nil {
		idempotencyKey = uuid.New()
		h.Deps.Logger.InfoContext(ctx, "idempotency key generated", slog.String("uuid", idempotencyKey.String()))
	} else {
		idempotencyKey = *req.Body.IdempotencyKey
	}

	res, err := h.TicketController.Create(ctx, models.TicketCreateIn{
		Queue:          req.Body.Queue,
		Title:          req.Body.Title,
		Description:    pointer.Get(req.Body.Description),
		IdempotencyKey: idempotencyKey,
	})
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &api.PostTickets201JSONResponse{
		CreatedAt:   res.CreatedAt,
		Description: res.Description,
		Key:         res.Key,
		Status:      api.TicketStatus(res.Status),
		Title:       res.Title,
		UpdatedAt:   res.UpdatedAt,
	}, nil
}

func (h *Handler) DeleteTicketsKey(
	ctx context.Context,
	req api.DeleteTicketsKeyRequestObject,
) (api.DeleteTicketsKeyResponseObject, error) {
	err := h.TicketController.Delete(ctx, req.Key)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return &api.DeleteTicketsKey404Response{}, nil
		}

		return nil, err
	}

	return &api.DeleteTicketsKey204Response{}, nil
}

func (h *Handler) GetTicketsKey(
	ctx context.Context,
	req api.GetTicketsKeyRequestObject,
) (api.GetTicketsKeyResponseObject, error) {
	ticket, err := h.TicketController.Get(ctx, req.Key)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return &api.GetTicketsKey404Response{}, nil
		}

		return nil, err
	}

	return &api.GetTicketsKey200JSONResponse{
		CreatedAt:   ticket.CreatedAt,
		Description: ticket.Description,
		Key:         ticket.Key,
		Status:      api.TicketStatus(ticket.Status),
		Title:       ticket.Title,
		UpdatedAt:   ticket.UpdatedAt,
	}, nil
}

func (h *Handler) PatchTicketsKey(
	ctx context.Context,
	req api.PatchTicketsKeyRequestObject,
) (api.PatchTicketsKeyResponseObject, error) {
	var status *string
	if req.Body.Status != nil {
		status = pointer.To(string(*req.Body.Status))
	}

	ticket, err := h.TicketController.Update(ctx, req.Key, models.TicketUpdateIn{
		Title:       req.Body.Title,
		Description: req.Body.Description,
		Status:      status,
	})
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return &api.PatchTicketsKey404Response{}, nil
		}

		return nil, err
	}

	return &api.PatchTicketsKey200JSONResponse{
		Description: pointer.To(ticket.Description),
		Status:      pointer.To(api.TicketPatchStatus(ticket.Status)),
		Title:       pointer.To(ticket.Title),
	}, nil
}
