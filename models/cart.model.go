package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	UserId uint       `json:"userId"`
	Total  float32    `json:"total"`
	Items  []CartItem `json:"items"`
}

// type CartItem struct {
// 	ID        uint    `json:"id" gorm:"primary_key"`
// 	UserId    uint    `json:"userId"`
// 	ProductId string  `json:"productId,omitempty" gorm:"type:varchar(100)"`
// 	Quantity  uint    `json:"quantity" gorm:"type:varchar(75)"`
// 	Product   Product `json:"product" gorm:"-:all"`
// }

type CartItem struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserId    uint               `json:"userId" bson:"userId"`
	ProductId primitive.ObjectID `json:"productId,omitempty" bson:"productId"`
	Quantity  uint               `json:"quantity" bson:"quantity"`
	Product   Product            `json:"product"`
	Products  []Product          `json:"products,omitempty"`
}
