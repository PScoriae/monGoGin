package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

var persons = []Person{{ID: 1, Name: "Pierre", Age: 20}, {ID: 2, Name: "Jonathan", Age: 30}, {ID: 3, Name: "Michael", Age: 25}}

func setupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")

	{
		users.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, persons)
		})

		users.POST("", func(c *gin.Context) {
			var json Person
			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			persons = append(persons, json)
			c.JSON(http.StatusOK, persons)
		})

		// Get users by ID
		users.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.String(http.StatusInternalServerError, "Internal Server Error")
			}
			for _, person := range persons {
				if person.ID == id {
					c.JSON(http.StatusOK, person)
					return
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"message": "No such person"})
		})
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
