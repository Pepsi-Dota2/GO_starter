package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	jwtKey := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["user_id"])

	return c.Next()
}
