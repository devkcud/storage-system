package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID           primitive.ObjectID `bson:"_id" json:"-"`
	Name         string             `bson:"name" json:"name"`
	Description  string             `bson:"description" json:"description"`
	PricePerUnit float64            `bson:"pricePerUnit" json:"pricePerUnit"`
	Quantity     int                `bson:"quantity" json:"quantity"`
	Tags         []string           `bson:"tags" json:"tags"`
}
