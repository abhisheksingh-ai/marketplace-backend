package models

import "time"

type Order struct {
	ID   string  `json:"id" bson:"_id,ommitempty"`
	ProducrID string `json:"product_id" bson:"product_id"`
	Quantity string `json:"quantity" bson:"quantity"`
	UserID string `json:"user_id" bson:"user_id"`
	OrderedAt time.Time `json:"ordered_at" bson:"ordered_at"`
}