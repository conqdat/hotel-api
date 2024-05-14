package main

import (
	"context"
	"log"

	"github.com/conqdat/hotel-api/api"
	"github.com/conqdat/hotel-api/db"
	_ "github.com/conqdat/hotel-api/docs" // Update import path to your docs package
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Hotel API
// @version 1.0
// @description This is a sample hotel API.
// @host localhost:3000
// @BasePath /api/v1

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New(config)

	// Swagger endpoint
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	apiV1 := app.Group("/api/v1")

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, db.DBNAME))
	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	apiV1.Put("/users/:id", userHandler.HandlePutUser)
	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Post("/users", userHandler.HandleCreateUser)

	app.Get("/", handleHelloWord)
	app.Listen(":3000")
}

// handleHelloWord godoc
// @Summary Show a Hello World message
// @Description Get a Hello World message
// @Tags root
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hello guys !!!"
// @Router / [get]
func handleHelloWord(c *fiber.Ctx) error {
	return c.JSON("Hello guys !!!")
}
