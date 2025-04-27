package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Price     string             `json:"price" bson:"price"`
	Category  string             `json:"category" bson:"category"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
