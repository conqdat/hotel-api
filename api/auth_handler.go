package api

import (
	"fmt"
	"github.com/conqdat/hotel-api/db"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userStore db.UserStore
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandle(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	user, err := h.userStore.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		fmt.Println("not found user")
		return fmt.Errorf("invalid credentials")
	}
	fmt.Println(user.EncryptedPassword)
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(params.Password))
	if err != nil {
		fmt.Println("fail to compare")
		return fmt.Errorf("invalid credentials")
	}
	fmt.Println("authenticated -> ", user)
	return nil
}
