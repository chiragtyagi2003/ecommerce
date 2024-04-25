package services

import (
	"github.com/Sebas3270/ecommerce-app-backend/db"
	"github.com/Sebas3270/ecommerce-app-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {

	// skipValue := c.QueryInt("offset", 0)

	// limitValue := c.QueryInt("limit", 15)

	// beginValue := c.Query("begin")
	// if beginValue != "" {
	// 	filter = bson.M{"name": bson.M{"$regex": beginValue, "$options": "im"}}
	// }

	var sales []models.Order

	err := db.MysqlConnection.Model(&models.Order{}).Preload("Items").Find(&sales).Error

	if err != nil {
		panic(err)
	}

	return c.Status(200).JSON(sales)
}
