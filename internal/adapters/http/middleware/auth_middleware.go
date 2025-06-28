package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/my-backend/internal/core/ports"
	"github.com/thitiphongD/my-backend/pkg/response"
)

// AuthMiddleware creates authentication middleware
func AuthMiddleware(authService ports.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, fiber.StatusUnauthorized, "Authorization header is required")
		}

		// Check if it starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return response.Error(c, fiber.StatusUnauthorized, "Authorization header must start with 'Bearer '")
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			return response.Error(c, fiber.StatusUnauthorized, "Token is required")
		}

		// Validate token
		user, err := authService.ValidateToken(token)
		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid or expired token")
		}

		// Store user ID in context
		c.Locals("userID", user.ID)
		c.Locals("user", user)

		return c.Next()
	}
}
