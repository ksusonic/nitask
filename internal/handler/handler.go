package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", Ping)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
