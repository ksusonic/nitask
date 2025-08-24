package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ksusonic/nitask/pkg/config"
)

type Server struct {
	*gin.Engine
}

func New(config config.ServerConfig) *Server {
	return &Server{
		Engine: gin.Default(),
	}
}
