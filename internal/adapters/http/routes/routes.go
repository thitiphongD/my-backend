package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/my-backend/internal/adapters/http/handlers"
	"github.com/thitiphongD/my-backend/internal/adapters/http/middleware"
	"github.com/thitiphongD/my-backend/internal/core/ports"
	"github.com/thitiphongD/my-backend/pkg/response"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, authService ports.AuthService, userService ports.UserService) {
	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// Health check route
	app.Get("/", func(c *fiber.Ctx) error {
		return response.Success(c, fiber.Map{
			"message": "Hello, World! From Daew project - Clean Architecture",
			"version": "v2.0.0",
		})
	})

	// Basic routes (demo purposes)
	app.Get("/say-hi/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name == "" {
			return response.Error(c, fiber.StatusBadRequest, "Name parameter is required")
		}
		return response.Success(c, fiber.Map{
			"message": "Hello, " + name,
		})
	})

	// JSON demo route
	app.Post("/json", func(c *fiber.Ctx) error {
		type RequestBody struct {
			Name string `json:"name" validate:"required"`
		}
		var body RequestBody
		if err := c.BodyParser(&body); err != nil {
			return response.Error(c, fiber.StatusBadRequest, "Invalid JSON")
		}
		if body.Name == "" {
			return response.Error(c, fiber.StatusBadRequest, "Name is required")
		}
		return response.Success(c, fiber.Map{
			"message": "JSON received successfully",
			"name":    body.Name,
		})
	})

	// API v1 routes
	v1 := app.Group("/api/v1")

	// Auth routes (public)
	auth := v1.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Get("/me", middleware.AuthMiddleware(authService), authHandler.GetMe)

	// User routes
	users := v1.Group("/users")
	users.Get("/", userHandler.GetUsers)                                                 // Public: Get all users
	users.Get("/:id", userHandler.GetUserByID)                                           // Public: Get user by ID
	users.Post("/", middleware.AuthMiddleware(authService), userHandler.CreateUser)      // Protected: Create user
	users.Put("/:id", middleware.AuthMiddleware(authService), userHandler.UpdateUser)    // Protected: Update user
	users.Delete("/:id", middleware.AuthMiddleware(authService), userHandler.DeleteUser) // Protected: Delete user
}
