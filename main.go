package main

import (
	"github.com/conqdat/hotel-api/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	apiV1 := app.Group("/api/v1")

	apiV1.Get("/users", api.HandleGetUser)
	app.Get("/", handleHelloWord)
	app.Listen(":3000")
}

func handleHelloWord(c *fiber.Ctx) error {
	return c.JSON("Hello guys !!!")
}
