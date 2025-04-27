package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/abhisheksingh-ai/marketplace-backend/internals/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProduct(mongoClient *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product

		//json data --> product struct
		if err := c.ShouldBindBodyWithJSON(&product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid request payload"})
			return
		}

		//time update
		product.CreatedAt = time.Now()

		//get collection
		productCollection := mongoClient.Database("marketplace").Collection("products")

		//creating a context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//insert into mongo
		_, err := productCollection.InsertOne(ctx, product)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error inserting new prodcut"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "new producr added successfully"})
	}
}
