package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/my-backend/internal/utils"
)

// AuthMiddleware checks for valid JWT token in Authorization header
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is required",
			})
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header must start with Bearer",
			})
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token is required",
			})
		}

		// Validate the token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Store user information in context for use in handlers
		c.Locals("userID", claims.UserID)
		c.Locals("userEmail", claims.Email)

		return c.Next()
	}
}

// OptionalAuthMiddleware checks for JWT token but doesn't require it
func OptionalAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString != "" {
				claims, err := utils.ValidateJWT(tokenString)
				if err == nil {
					c.Locals("userID", claims.UserID)
					c.Locals("userEmail", claims.Email)
				}
			}
		}
		return c.Next()
	}
}
