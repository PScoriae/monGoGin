package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingRoutes(router *gin.Engine) {

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
