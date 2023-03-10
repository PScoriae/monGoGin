package handlers

import (
	"context"
	"mongogin/internal/app/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	// prefill user object
	user := User{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// validation
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insert validated json
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

func GetAllUsers(c *gin.Context) {
	// get client
	client, err := db.GetMongoClient()
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// get cursor
	cursor, err := client.Database("mongogin-prod").Collection(string(db.Users)).Find(context.TODO(), bson.D{})
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// decodes cursor into results
	var results []User
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// returns values
	c.JSON(http.StatusOK, results)
}
