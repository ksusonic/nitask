package handler

import (
	"context"

	"github.com/ksusonic/nitask/pkg/api"
)

func (h *Handler) GetPing(ctx context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {
	return api.GetPing200JSONResponse{Ping: "pong"}, nil
}
