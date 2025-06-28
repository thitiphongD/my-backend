package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/thitiphongD/my-backend/internal/adapters/database/repositories"
	"github.com/thitiphongD/my-backend/internal/adapters/http/routes"
	"github.com/thitiphongD/my-backend/internal/config"
	"github.com/thitiphongD/my-backend/internal/core/domain"
	"github.com/thitiphongD/my-backend/internal/core/services"
	"github.com/thitiphongD/my-backend/internal/database"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	database.ConnectDatabase()
	db := database.GetDB()

	// Auto migrate the schema
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services with dependency injection
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		},
	})

	// Global middlewares
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${status} - ${method} ${path} - ${latency}\n",
	}))

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:5173,http://127.0.0.1:3000,http://127.0.0.1:5173",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(app, authService, userService)

	// Start server
	port := ":" + cfg.Port
	log.Printf("🚀 Server starting on port %s", cfg.Port)
	log.Printf("📚 API Documentation available at http://localhost%s", port)
	log.Printf("🏥 Health check at http://localhost%s/", port)

	if err := app.Listen(port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
