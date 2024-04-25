package models

import "time"

type Order struct {
	ID     uint        `json:"id" gorm:"primary_key"`
	UserId uint        `json:"userId"`
	Total  float32     `json:"total" gorm:"type:varchar(75)"`
	Date   time.Time   `json:"date" gorm:"type:varchar(75)"`
	Status string      `json:"status" gorm:"type:varchar(50)"`
	Items  []OrderItem `json:"items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	OrderId   uint    `json:"orderId"`
	ProductId string  `json:"productId" gorm:"type:varchar(100)"`
	Quantity  uint    `json:"quantity" gorm:"type:varchar(75)"`
	Price     float32 `json:"price" gorm:"type:varchar(75)"`
}
