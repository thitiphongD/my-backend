package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/thitiphongD/my-backend/internal/auth"
	"github.com/thitiphongD/my-backend/internal/database"
	"github.com/thitiphongD/my-backend/internal/middleware"
	"github.com/thitiphongD/my-backend/internal/models"
)

func main() {
	// Load .env file if it exists (for development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database connection
	database.ConnectDatabase()

	// Auto migrate the schema
	err := database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Initialize auth service
	authService := auth.NewAuthService()

	app := fiber.New()

	// Global middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:5173,http://127.0.0.1:3000,http://127.0.0.1:5173", // React/Vite development servers
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
	}))

	// Public routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! From Daew project")
	})

	app.Get("/say-hi/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Name parameter is required")
		}
		return c.SendString("Hello, " + name)
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		if name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Name is required")
		}
		return c.SendString("Form submitted successfully with name: " + name)
	})

	app.Post("/json", func(c *fiber.Ctx) error {
		type RequestBody struct {
			Name string `json:"name"`
		}
		var body RequestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
		}
		if body.Name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Name is required")
		}
		return c.JSON(fiber.Map{
			"message": "JSON received successfully",
			"name":    body.Name,
		})
	})

	// Auth routes (public)
	app.Post("/auth/register", func(c *fiber.Ctx) error {
		var request models.RegisterRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		if request.Name == "" || request.Email == "" || request.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Name, email, and password are required",
			})
		}

		response, err := authService.Register(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	})

	app.Post("/auth/login", func(c *fiber.Ctx) error {
		var request models.LoginRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		if request.Email == "" || request.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email and password are required",
			})
		}

		response, err := authService.Login(request)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(response)
	})

	// Protected routes (require authentication)
	app.Get("/auth/me", middleware.AuthMiddleware(), func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)
		user, err := authService.GetUserByID(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get user",
			})
		}
		return c.JSON(user)
	})

	// Public database API endpoints
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User
		result := database.GetDB().Find(&users)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch users",
			})
		}
		return c.JSON(users)
	})

	// Protected endpoint - Create user (admin only)
	app.Post("/users", middleware.AuthMiddleware(), func(c *fiber.Ctx) error {
		var request models.CreateUserRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		if request.Name == "" || request.Email == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Name and email are required",
			})
		}

		user := models.User{
			Name:  request.Name,
			Email: request.Email,
		}

		result := database.GetDB().Create(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user models.User
		result := database.GetDB().First(&user, id)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.JSON(user)
	})

	// Health check endpoint for database
	app.Get("/health", func(c *fiber.Ctx) error {
		sqlDB, err := database.GetDB().DB()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":   "error",
				"database": "disconnected",
			})
		}

		if err := sqlDB.Ping(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":   "error",
				"database": "ping failed",
			})
		}

		return c.JSON(fiber.Map{
			"status":   "ok",
			"database": "connected",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set in the .env file")
		port = "8080"
	}
	log.Println("Server is running on port " + port)

	app.Listen(":" + port)
}
