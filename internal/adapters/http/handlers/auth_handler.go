package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/my-backend/internal/core/domain"
	"github.com/thitiphongD/my-backend/internal/core/ports"
	"github.com/thitiphongD/my-backend/pkg/response"
	"github.com/thitiphongD/my-backend/pkg/validator"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	authService ports.AuthService
}

// NewAuthHandler creates a new auth handler instance
func NewAuthHandler(authService ports.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req domain.RegisterRequest

	if err := validator.ParseAndValidate(c, &req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	authResponse, err := h.authService.Register(&req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Created(c, authResponse, "User registered successfully")
}

// Login handles user login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req domain.LoginRequest

	if err := validator.ParseAndValidate(c, &req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	authResponse, err := h.authService.Login(&req)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	return response.Success(c, authResponse, "Login successful")
}

// GetMe returns the current authenticated user's information
func (h *AuthHandler) GetMe(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return response.Error(c, fiber.StatusUnauthorized, "User not authenticated")
	}

	user, err := h.authService.GetUserByID(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, user, "User information retrieved successfully")
}
