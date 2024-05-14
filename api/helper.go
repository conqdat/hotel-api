package api

import (
	"fmt"
	"github.com/conqdat/hotel-api/types"
	"github.com/gofiber/fiber/v2"
)

func getAuthUser(c *fiber.Ctx) (*types.User, error) {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return nil, fmt.Errorf("failed to get user from context")
	}
	return user, nil
}
