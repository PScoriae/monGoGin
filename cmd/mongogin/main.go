package main

import (
	"mongogin/internal/app/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")

	{
		users.GET("", handlers.GetAllUsers)

		users.POST("", handlers.CreateUser)

		// // Get users by ID
		// users.GET("/:id", func(c *gin.Context) {
		// 	id, err := strconv.Atoi(c.Param("id"))
		// 	if err != nil {
		// 		c.String(http.StatusInternalServerError, "Internal Server Error")
		// 	}
		// 	for _, person := range persons {
		// 		if person.ID == id {
		// 			c.JSON(http.StatusOK, person)
		// 			return
		// 		}
		// 	}
		// 	c.JSON(http.StatusNotFound, gin.H{"message": "No such person"})
		// })
	}
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
