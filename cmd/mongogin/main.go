package main

import (
	"mongogin/internal/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.AddRoutes(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
