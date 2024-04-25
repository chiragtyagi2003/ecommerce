package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection *mongo.Client
var ProductsCollection *mongo.Collection
var CartCollection *mongo.Collection

func ConnectMongoDb() {

	err := godotenv.Load()

	if err != nil {
		print("Error loading .env file")
	}

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGO_CONNECTION")).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	MongoConnection = client
	ProductsCollection = MongoConnection.Database("ecommerce_db").Collection("products")
	CartCollection = MongoConnection.Database("ecommerce_db").Collection("cart_items")

}
