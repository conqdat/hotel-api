package main

import (
	"context"
	"github.com/conqdat/hotel-api/api"
	"github.com/conqdat/hotel-api/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userCollections = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiV1.Get("/user/:id", userHandler.HandleGetUser)

	app.Get("/", handleHelloWord)
	app.Listen(":3000")
}

func handleHelloWord(c *fiber.Ctx) error {
	return c.JSON("Hello guys !!!")
}
