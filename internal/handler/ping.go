package handler

import (
	"context"

	"github.com/ksusonic/nitask/pkg/api"
)

func (h *Handler) GetPing(context.Context, api.GetPingRequestObject) (api.GetPingResponseObject, error) {
	return api.GetPing200JSONResponse{Ping: "pong"}, nil
}
