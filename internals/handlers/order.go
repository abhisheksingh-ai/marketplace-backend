package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abhisheksingh-ai/marketplace-backend/internals/models"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateOrder(mongoClient *mongo.Client, rabbitChannel *amqp.Channel) gin.HandlerFunc {
	return func(c *gin.Context) {

		var order models.Order

		//Binding incoming json to order struct
		if err := c.ShouldBindBodyWithJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		//set order time
		order.OrderedAt = time.Now()

		//marshal order info order --> byte
		orderBytes, err := json.Marshal(order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error serializing order"})
			return
		}

		//publish to rabbitmq
		err = rabbitChannel.Publish(
			"",
			"orders",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        orderBytes,
			},
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed tp publish the order"})
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order Placed Succesfull"})
	}
}
