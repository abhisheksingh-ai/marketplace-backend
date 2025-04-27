package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/abhisheksingh-ai/marketplace-backend/internals/handlers"
	"github.com/abhisheksingh-ai/marketplace-backend/internals/utils"
)

func main() {
	//Load enviroment varibles from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in go")
	}

	//connect to mongodb
	mongoClient, err := utils.ConnectMongo()
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}
	defer mongoClient.Disconnect(nil)

	//connect to rabbitmq
	rabbitConn, rabbitChannel, err := utils.ConnectRabbitMQ()

	if err != nil {
		log.Fatal("RabbitMQ connection error: ", err)
	}

	defer rabbitConn.Close()
	defer rabbitChannel.Close()

	//start RabbitMQ consumer in background
	go utils.StartOrderConsumer(rabbitChannel, mongoClient)

	//Initializing gin router
	router := gin.Default()

	// Register routes
	router.POST("/users", handlers.CreateUser(mongoClient))
	// router.POST("/products", handlers.CreateProduct(mongoClient))
	router.POST("/orders", handlers.CreateOrder(mongoClient, rabbitChannel))

	// Start server on given port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running at port", port)
	router.Run(":" + port)
}
