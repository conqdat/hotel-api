package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/conqdat/hotel-api/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok {
			return fmt.Errorf("unauthorized")
		}
		claims, err := validateToken(token[0])
		if err != nil {
			return Unauthorized()
		}

		// Check if the "expires" field exists in the claims map
		expiresFloat, ok := claims["expires"].(float64)
		if !ok {
			return NewError(http.StatusUnauthorized, "expiration time not found or not a valid time")
		}

		expires := int64(expiresFloat)
		// Check token expiration
		if time.Now().Unix() > expires {
			return fmt.Errorf("token expired")
		}
		userID := claims["id"].(string)
		user, err := userStore.GetUserByID(c.Context(), userID)
		if err != nil {
			return NewError(http.StatusUnauthorized, "Unauthorized")
		}
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, NewError(http.StatusUnauthorized, "Unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, NewError(http.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, NewError(http.StatusUnauthorized, "Unauthorized")
	}
	return claims, nil
}
