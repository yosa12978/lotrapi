package api

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/lotrapi/internal/api/middleware"
)

func Run(listener net.Listener) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.RequestLogging())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello lotr!",
		})
	})
	router.RunListener(listener)
}
