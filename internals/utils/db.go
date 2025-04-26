package utils

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, error) {
	//get mongouri from .env
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env")
	}

	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//connectiong to mongodb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	// verifying mongoConnection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("MongoDB connected Succesfully")
	return client, nil
}
