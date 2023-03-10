package handlers

import (
	"context"
	"mongogin/internal/app/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Age       int                `json:"age" bson:"age" binding:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at"`
}

func CreateUser(c *gin.Context) {
	// validation
	user := User{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := db.GetMongoClient()

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	_, err = client.Database("mongogin-prod").Collection(string(db.Users)).InsertOne(context.TODO(), user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusCreated, user)
}
