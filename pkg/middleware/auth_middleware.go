package middleware

import (
	"github.com/chicken-afk/go-fiber/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	claims, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	if id, ok := claims["id"].(float64); ok && id != 83 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "failed",
			"message": "Do not enter this area",
		})
	} else {
		return c.Next()
	}

}
