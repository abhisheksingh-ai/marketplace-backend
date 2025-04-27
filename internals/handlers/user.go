package handlers

import (
	"net/http"
	"time"

	"github.com/abhisheksingh-ai/marketplace-backend/internals/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

func CreateUser(mongoCLient *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		//Bind json to user struct
		if err := c.ShouldBindBodyWithJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		//set createdAt field
		user.CreatedAt = time.Now()

		//get user collection
		userCollection := mongoCLient.Database("marketplace").Collection("users")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//inser into mongodb
		_, err := userCollection.InsertOne(ctx, user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user has been created"})
	}
}
