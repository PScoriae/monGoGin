package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var persons = []Person{{ID: 1, Name: "Pierre", Age: 20}, {ID: 2, Name: "Jonathan", Age: 30}, {ID: 3, Name: "Michael", Age: 25}}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get users
	r.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, persons)
	})

	// Get users by ID
	r.GET("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		for _, person := range persons {
			if person.ID == id {
				c.IndentedJSON(http.StatusOK, person)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such person"})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
