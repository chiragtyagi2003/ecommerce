package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" validate:"omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"omitempty"`
	Price       float64            `json:"price" bson:"price" validate:"required" `
	Image       string             `json:"image" bson:"image" validate:"omitempty" `
	Tags        []string           `json:"tags" bson:"tags" validate:"omitempty"`
}
