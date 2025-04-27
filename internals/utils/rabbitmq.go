package utils

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/abhisheksingh-ai/marketplace-backend/internals/models"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

// This function will connect to rabbitmq server
func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {

	rabbitURL := os.Getenv("RABBITMQ_URL")

	if rabbitURL == "" {
		log.Fatal("RABBITMQ_URL is not set in dot env")
	}

	conn, err := amqp.DialConfig(rabbitURL, amqp.Config{})

	if err != nil {
		return nil, nil, err
	}

	channel, err := conn.Channel()

	if err != nil {
		return nil, nil, err
	}

	log.Println("RabbitMQ connection and channel made successfully")

	return conn, channel, nil
}

// StartOrderConsumer listen to "order" queue and saves orders into mongo db

func StartOrderConsumer(channel *amqp.Channel, mongoClient *mongo.Client) {
	//ensure queue exists
	queue, err := channel.QueueDeclare(
		"orders",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Queue declare error:", err)
	}

	//consume message
	messages, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Queue Consume Error:", err)
	}

	//mongodb collection for order
	orderCollection := mongoClient.Database("marketplace").Collection("orders")

	log.Println("Waiting for new orders from rabbitmq...")

	//start listening
	for msg := range messages {

		var order models.Order //by this i am creating a type of documet that will inserted in mongodb as document

		//decode json messgae
		err := json.Unmarshal(msg.Body, &order)

		if err != nil {
			log.Fatal("Decoding message error", err)
			continue
		}

		//inserting into mongodb
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err = orderCollection.InsertOne(ctx, order)

		if err != nil {
			log.Println("Error inserting order:", err)
			continue
		}

		log.Println("Order saved to mongodb:", order)
	}
}
