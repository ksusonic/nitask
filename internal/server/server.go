package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/nitask/internal/server/handler"
	"github.com/ksusonic/nitask/pkg/api"
	"github.com/ksusonic/nitask/pkg/config"
)

const (
	ReadHeaderTimeout = 5 * time.Second
)

type Server struct {
	*gin.Engine
}

func New(
	cfg config.ServerConfig,
	handlerDeps *handler.Deps,
) *Server {
	r := gin.Default()

	gin.SetMode(cfg.Mode)

	server := api.NewStrictHandler(handler.New(handlerDeps), nil)

	api.RegisterHandlers(r, server)

	return &Server{
		Engine: r,
	}
}
