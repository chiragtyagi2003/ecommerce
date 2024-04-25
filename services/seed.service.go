package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Sebas3270/ecommerce-app-backend/db"
	"github.com/Sebas3270/ecommerce-app-backend/helpers"
	"github.com/Sebas3270/ecommerce-app-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RunSeed(c *fiber.Ctx) error {

	db.ProductsCollection.Drop(context.TODO())
	db.CartCollection.Drop(context.TODO())
	db.MysqlConnection.Exec("DELETE FROM user")

	products := getProductsSeed()
	db.ProductsCollection.InsertMany(context.TODO(), products)

	users := getUserSeed()
	db.MysqlConnection.Create(&users)

	return c.Status(200).SendString("Seed created successfully")

}

func getProductsSeed() []interface{} {

	var productsJson string
	var products []models.Product

	productsJson = `
	[
		{
			"name": "PS4 Slim",
			"description": "A PS4 Console with 256 gb capacity",
			"price": 5999.99,
			"image": "https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/PS4-Console-wDS4.png/2560px-PS4-Console-wDS4.png",
			"tags": [
				"technology",
				"videogames"
			]
		},
		{
			"name": "Air Pods Pro",
			"description": "The best bluetooth headphones by Samsung",
			"price": 20000,
			"image": "https://assets.stickpng.com/images/60b79e8771a1fd000411f6be.png",
			"tags": [
				"technology",
				"headphones"
			]
		},
		{
			"name": "Huawei P50",
			"description": "A Huawei cellphone that is amazing, The simplified, geometric design makes the Dual-Matrix Camera Design truly stand out. The two \"dazzling eyes\" shine through, wherever you go.",
			"price": 2547.3,
			"image": "https://img01.huaweifile.com/sg/ms/co/pms/uomcdn/CO_HW_B2C/pms/202208/gbom/6941487261024/428_428_B787FD155916721059A1EED309C9715Bmp.png",
			"tags": [
				"technology",
				"cellphones"
			]
		},
		{
			"name": "Galaxy Buds",
			"description": "The best bluetooth headphones by Samsung",
			"price": 2296.2,
			"image": "https://cdn.shopify.com/s/files/1/0805/4543/products/SamsungGalaxyBluetoothBuds-Black_1024x1024.png?v=1600331876",
			"tags": [
				"technology",
				"headphones"
			]
		},
		{
			"name": "Huawei P40 pro",
			"description": "The best Huawei cellphone by Huawei",
			"price": 12999.99,
			"image": "https://assets.stickpng.com/images/61d4a4168b51e20004664d49.png",
			"tags": [
				"technology",
				"cellphones"
			]
		},
		{
			"name": "Xbox Series S",
			"description": "A Xbox Series S Console with 256 gb capacity",
			"price": 4999.99,
			"image": "https://assets.xboxservices.com/assets/ea/f2/eaf2bd1e-f245-448b-891b-b276b32a8a24.png?n=294022_BuyBox01_Image-D.png",
			"tags": [
				"technology",
				"videogames"
			]
		},
		{
			"name": "iPhone 13",
			"description": "An amazing iphone to do all you want to do",
			"price": 14999.99,
			"image": "https://png.monster/wp-content/uploads/2022/09/png.monster-210.png",
			"tags": [
				"technology",
				"cellphones"
			]
		},
		{
			"name": "Five Nights At Freddy's Xbox Game",
			"description": "The best horror videogame",
			"price": 39.99,
			"image": "https://xboxofflinechile.com/wp-content/uploads/2022/04/Disen%CC%83o-sin-ti%CC%81tulo.png",
			"tags": [
				"technology",
				"videogames"
			]
		}
	]
	`

	err := json.Unmarshal([]byte(productsJson), &products)

	if err != nil {
		log.Println(err)
	}

	newValue := make([]interface{}, len(products))

	for i := range products {
		products[i].Id = primitive.NewObjectID()
		newValue[i] = products[i]
	}

	return newValue

}

func getUserSeed() []models.User {
	var usersJson string
	// users := new([]models.User)
	var users []models.User

	usersJson = `
	[
		{
			"firstName": "Sebastián",
			"lastName": "Álvarez Ocampo",
			"email": "test@gmail.com",
			"password": "password",
			"address": "UNIVA 2745 Av. Tepeyac",
			"telephone": "3344225566",
			"total": 0
		},
		{
			"firstName": "David",
			"lastName": "Silva Romo",
			"email": "test2@gmail.com",
			"password": "password",
			"address": "Apple 22 Av. Crystal",
			"telephone": "3344225566",
			"total": 0
		}
	]
	`

	err := json.Unmarshal([]byte(usersJson), &users)

	if err != nil {
		log.Println(err)
	}

	for _, user := range users {
		// log.Print(*user.Password)
		pass, _ := helpers.HashPassword(*user.Password)
		*user.Password = pass
	}

	return users

}
