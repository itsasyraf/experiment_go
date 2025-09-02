package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing bearer token"})
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := ParseToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		// store claims in context for handlers
		c.Locals("claims", claims)
		return c.Next()
	}
}

// Helper to read claims in handlers
func GetClaims(c *fiber.Ctx) *UserClaims {
	if v := c.Locals("claims"); v != nil {
		if claims, ok := v.(*UserClaims); ok {
			return claims
		}
	}
	return nil
}
