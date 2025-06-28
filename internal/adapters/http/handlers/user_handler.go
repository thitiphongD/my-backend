package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/my-backend/internal/core/domain"
	"github.com/thitiphongD/my-backend/internal/core/ports"
	"github.com/thitiphongD/my-backend/pkg/response"
	"github.com/thitiphongD/my-backend/pkg/validator"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userService ports.UserService
}

// NewUserHandler creates a new user handler instance
func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser handles user creation
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req domain.CreateUserRequest

	if err := validator.ParseAndValidate(c, &req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Created(c, user, "User created successfully")
}

// GetUsers handles retrieving all users
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, users, "Users retrieved successfully")
}

// GetUserByID handles retrieving a user by ID
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}

	return response.Success(c, user, "User retrieved successfully")
}

// UpdateUser handles user updates
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req domain.CreateUserRequest
	if err := validator.ParseAndValidate(c, &req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := h.userService.UpdateUser(uint(id), &req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, user, "User updated successfully")
}

// DeleteUser handles user deletion
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, nil, "User deleted successfully")
}
