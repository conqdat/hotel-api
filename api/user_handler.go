package api

import (
	"github.com/conqdat/hotel-api/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(ctx *fiber.Ctx) error {
	user := types.User{
		FirstName: "Tran",
		LastName:  "Dat",
	}
	return ctx.JSON(user)
}
