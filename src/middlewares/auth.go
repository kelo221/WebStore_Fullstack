package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func IsAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	return c.Next()
}

func GetUserID(c *fiber.Ctx) (string, error) {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err

	}

	payload := token.Claims.(*jwt.StandardClaims)

	return payload.Subject, nil

}
