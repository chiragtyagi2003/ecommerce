package main

import (
	"github.com/Sebas3270/ecommerce-app-backend/db"
	"github.com/Sebas3270/ecommerce-app-backend/services"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()

	db.ConnectMysqlDb()
	db.ConnectMongoDb()

	api := app.Group("/api")

	//Seed
	api.Get("/seed", services.RunSeed)

	// Group routes
	authRoute := api.Group("/auth")
	productsRoute := api.Group("/products")
	cartRoute := api.Group("/carts")
	salesRoute := api.Group("/sales")

	// Public routes

	// Auth Routes
	authRoute.Post("/login", services.LogIn)
	authRoute.Post("/register", services.Resgister)

	//Product Routes
	productsRoute.Get("/", services.GetProducts)

	//JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
		ContextKey: "auth",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		},
	}))

	//Private routes

	//Auth Routes
	authRoute.Post("/renew", services.RenewToken)

	//Cart routes
	cartRoute.Get("/", services.GetCarts)
	cartRoute.Post("/", services.AddCartItem)
	cartRoute.Delete("/", services.DeleteCartItem)
	cartRoute.Put("/", services.UpdateCartItem)

	//Sales routes
	salesRoute.Get("/", services.GetOrders)

	app.Listen(":3000")

	// defer func() {
	// 	dbInstance, _ := db.MysqlConnection.DB()
	// 	_ = dbInstance.Close()

	// 	db.MongoConnection.Disconnect(context.TODO())

	// 	log.Print("Closed connections")
	// }()
}
