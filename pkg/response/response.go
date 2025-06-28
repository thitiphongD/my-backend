package response

import "github.com/gofiber/fiber/v2"

// APIResponse represents a standard API response structure
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Success returns a successful response
func Success(c *fiber.Ctx, data interface{}, message ...string) error {
	response := APIResponse{
		Success: true,
		Data:    data,
	}

	if len(message) > 0 {
		response.Message = message[0]
	}

	return c.JSON(response)
}

// Error returns an error response
func Error(c *fiber.Ctx, statusCode int, error interface{}, message ...string) error {
	response := APIResponse{
		Success: false,
		Error:   error,
	}

	if len(message) > 0 {
		response.Message = message[0]
	}

	return c.Status(statusCode).JSON(response)
}

// Created returns a created response (201)
func Created(c *fiber.Ctx, data interface{}, message ...string) error {
	response := APIResponse{
		Success: true,
		Data:    data,
	}

	if len(message) > 0 {
		response.Message = message[0]
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
