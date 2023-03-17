package routes

import (
	"mongogin/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func userRoutes(routerGroup *gin.Engine) {

	users := routerGroup.Group("/users")

	{
		users.GET("", handlers.GetAllUsers)

		users.POST("", handlers.CreateUser)

		users.GET("/:id", handlers.GetUsersById)
	}
}
