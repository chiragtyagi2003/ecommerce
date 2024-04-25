package services

import (
	"context"
	"log"

	"github.com/Sebas3270/ecommerce-app-backend/db"
	"github.com/Sebas3270/ecommerce-app-backend/helpers"
	"github.com/Sebas3270/ecommerce-app-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func GetCarts(c *fiber.Ctx) error {

// 	userId := helpers.GetTokenData(c)

// 	var user models.User
// 	db.MysqlConnection.Where("id = ?", userId).Preload("Items").First(&user)

// 	// var carItems []models.CartItem
// 	// db.MysqlConnection.Where("user_id = ?", userId).Find(&carItems)

// 	var cart models.Cart
// 	cart.UserId = user.ID
// 	cart.Items = *user.Items
// 	cart.Total = user.CartTotal

// 	if len(cart.Items) == 0 {
// 		user.CartTotal = 0
// 		go db.MysqlConnection.Save(&user)
// 		return c.Status(200).JSON(cart)
// 	}

// 	var ids []string
// 	for _, cartItem := range cart.Items {
// 		ids = append(ids, cartItem.ProductId)
// 	}

// 	products, err := GetProductsByIds(c, ids)

// 	if err != nil {
// 		panic(err)
// 	}

// 	cart.Total = 0

// 	for index := range cart.Items {
// 		price := (float32(products[index].Price) * float32(cart.Items[index].Quantity))
// 		cart.Items[index].Product = products[index]
// 		cart.Total = cart.Total + price
// 		log.Print(price)
// 	}

// 	user.CartTotal = cart.Total
// 	go db.MysqlConnection.Save(&user)

// 	return c.Status(200).JSON(cart)
// }

func GetCarts(c *fiber.Ctx) error {

	userId := helpers.GetTokenData(c)

	var user models.User
	err := db.MysqlConnection.Where("id = ?", userId).First(&user).Error

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	var cart models.Cart
	cart.UserId = user.ID
	cart.Total = user.CartTotal

	aggPopulate := bson.M{"$lookup": bson.M{
		"from":         "products",  // the collection name
		"localField":   "productId", // the field on the child struct
		"foreignField": "_id",       // the field on the parent struct
		"as":           "products",  // the field to populate into
	}}

	searchById := bson.M{"$match": bson.M{"userId": user.ID}}

	cursor, err := db.CartCollection.Aggregate(context.TODO(), []bson.M{
		searchById, aggPopulate,
	})

	if err != nil {
		log.Print("Error here")
		return err
	}

	var items []models.CartItem

	if err = cursor.All(context.TODO(), &items); err != nil {
		log.Print("Error here")
		return err
	}

	// var total float32
	cart.Total = 0

	if items == nil || len(items) == 0 {
		log.Print("Empty items")
		db.MysqlConnection.Model(&user).Update("cart_total", 0)
		return c.JSON(&cart)
	}

	cart.Items = items

	for index := range cart.Items {
		cart.Items[index].Product = items[index].Products[0]
		cart.Items[index].Products = nil
		cart.Total += float32(cart.Items[index].Product.Price) * (float32(cart.Items[index].Quantity))
	}

	// cart.Total = total
	user.CartTotal = cart.Total

	go db.MysqlConnection.Model(&user).Update("cart_total", cart.Total)

	return c.JSON(&cart)
}

func AddCartItem(c *fiber.Ctx) error {

	newCartItem := new(models.CartItem)
	if err := c.BodyParser(newCartItem); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	userId := helpers.GetTokenData(c)
	newCartItem.UserId = userId

	// db.MysqlConnection.Create(&newCartItem)
	newCartItem.ID = primitive.NewObjectID()
	newCartItem.Quantity = 1
	_, err := db.CartCollection.InsertOne(context.TODO(), newCartItem)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Error")
	}

	return c.Status(200).JSON(&newCartItem)
}

func DeleteCartItem(c *fiber.Ctx) error {

	userId := helpers.GetTokenData(c)

	cartItem := new(models.CartItem)

	if err := c.BodyParser(cartItem); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// err := db.MysqlConnection.Where("id = ?", cartItem.ID).First(&cartItem).Error
	filter := bson.D{{Key: "_id", Value: cartItem.ID}}
	err := db.CartCollection.FindOne(context.TODO(), filter).Decode(&cartItem)

	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	}

	if cartItem.UserId != userId {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not allowed to delete this resource")
	}

	// go db.MysqlConnection.Delete(&cartItem)
	go db.CartCollection.DeleteOne(context.TODO(), filter)

	return c.Status(200).SendString("Product deleted successfully")
}

func UpdateCartItem(c *fiber.Ctx) error {

	userId := helpers.GetTokenData(c)
	cartItem := new(models.CartItem)

	if err := c.BodyParser(cartItem); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// err := db.MysqlConnection.Where("id = ?", cartItem.ID).First(&cartItem).Error
	filter := bson.D{{Key: "_id", Value: cartItem.ID}}
	err := db.CartCollection.FindOne(context.TODO(), filter).Decode(&cartItem)

	if err != nil {
		return c.Status(400).SendString("Item not found")
	}

	if cartItem.UserId != userId {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not allowed to update this resource")
	}

	action := c.Query("action", "add")

	if action == "add" {
		cartItem.Quantity++
	}

	if action == "reduce" {
		cartItem.Quantity--
	}

	// go db.MysqlConnection.Save(&cartItem)

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "quantity", Value: cartItem.Quantity}}}}
	go db.CartCollection.UpdateOne(context.TODO(), filter, update)

	return c.Status(200).SendString("Product updated successfully")

}
