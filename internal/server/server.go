package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/nitask/pkg/api"
)

const (
	ReadHeaderTimeout = 5 * time.Second
)

type Server struct {
	*gin.Engine
}

func New(impl api.StrictServerInterface) *Server {
	r := gin.Default()

	server := api.NewStrictHandler(impl, nil)

	api.RegisterHandlers(r, server)

	return &Server{
		Engine: r,
	}
}
