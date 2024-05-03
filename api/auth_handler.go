package api

import (
	"fmt"
	"os"
	"time"

	"github.com/conqdat/hotel-api/db"
	"github.com/conqdat/hotel-api/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	userStore db.UserStore
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User *types.User	`json:"user"`
	Token string 		`json:"token"`
}

type genericResp struct {
	Type    string `json:"type"`
	Message string `json:"message"`
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
		return fmt.Errorf("invalid credentials")
	}
	if !types.IsValidPassword(user.EncryptedPassword, params.Password) {
		return fmt.Errorf("invalid credentials")
	}
	token := createTokenFromUser(user)
	userRes := AuthResponse{
		User: user,
		Token: token,
	}
	return c.JSON(userRes)
}

func createTokenFromUser(user *types.User) string {
	now := time.Now()
	expires := now.Add(time.Hour * 4)
	claims := jwt.MapClaims{
		"id":        user.ID,
		"email":     user.Email,
		"expires": expires.Unix(), 
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
	secret := os.Getenv("JWT_SECRET")
	tokenSt, err := token.SignedString([]byte(secret)) 
	if err != nil {
		return "fail to generate token"
	}
	return tokenSt
}
