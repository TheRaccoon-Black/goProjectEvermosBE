package http

import (
	"goProjectEvermos/pkg/helper"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return ErrorResponse(c, fiber.StatusUnauthorized, "Header otorisasi tidak ditemukan", nil)
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return ErrorResponse(c, fiber.StatusUnauthorized, "Format token tidak valid", nil)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := helper.ValidateToken(tokenString)
		if err != nil {
			return ErrorResponse(c, fiber.StatusUnauthorized, "Token tidak valid", err.Error())
		}

		c.Locals("userID", claims.UserID)
		c.Locals("userRole", claims.Role)

		return c.Next()
	}
}