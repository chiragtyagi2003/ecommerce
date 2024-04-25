package services

import (
	"context"
	"log"

	"github.com/Sebas3270/ecommerce-app-backend/db"
	"github.com/Sebas3270/ecommerce-app-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProducts(c *fiber.Ctx) error {

	skipValue := c.QueryInt("offset", 0)
	limitValue := c.QueryInt("limit", 15)
	beginValue := c.Query("begin")
	tagValue := c.Query("tag")

	var filter primitive.D

	if tagValue != "" {
		filter = bson.D{
			{Key: "tags", Value: bson.D{{Key: "$all", Value: bson.A{tagValue}}}},
			{Key: "name", Value: bson.M{"$regex": beginValue, "$options": "im"}},
		}
	} else {
		filter = bson.D{
			{Key: "name", Value: bson.M{"$regex": beginValue, "$options": "im"}},
		}
	}

	options := options.Find()
	options.SetSkip(int64(skipValue))
	options.SetLimit(int64(limitValue))

	var products []models.Product

	cursor, err := db.ProductsCollection.Find(context.TODO(), filter, options)

	if err != nil {
		log.Print(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error": "Error getting products",
		})
	}

	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Print(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"error": "Error getting products",
		})
	}

	return c.Status(200).JSON(products)
}

func transformIds(ids *[]string, mongoIds *[]primitive.ObjectID) {

	for _, id := range *ids {
		// element is the element from someSlice for where we are
		idCreated, _ := primitive.ObjectIDFromHex(id)
		*mongoIds = append(*mongoIds, idCreated)
	}

}

func GetProductsByIds(c *fiber.Ctx, ids []string) ([]models.Product, error) {

	var mongoIds []primitive.ObjectID

	transformIds(&ids, &mongoIds)

	var filter primitive.M

	filter = bson.M{"_id": bson.M{"$in": mongoIds}}

	var products []models.Product

	cursor, err := db.ProductsCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return products, nil
}

func CreateProduct(c *fiber.Ctx) error {

	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	product.Id = primitive.NewObjectID()

	_, err := db.ProductsCollection.InsertOne(context.TODO(), product)

	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": "Error creating products",
		})
	}

	return c.JSON(product)
}
