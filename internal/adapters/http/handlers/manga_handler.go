package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/my-backend/internal/core/domain"
	"github.com/thitiphongD/my-backend/internal/core/ports"
	"github.com/thitiphongD/my-backend/pkg/response"
	"github.com/thitiphongD/my-backend/pkg/validator"
)

// MangaHandler handles HTTP requests for manga operations
type MangaHandler struct {
	mangaService ports.MangaService
}

// NewMangaHandler creates a new manga handler instance
func NewMangaHandler(mangaService ports.MangaService) *MangaHandler {
	return &MangaHandler{
		mangaService: mangaService,
	}
}

// CreateManga handles POST /api/v1/mangas
func (h *MangaHandler) CreateManga(c *fiber.Ctx) error {
	var req domain.CreateMangaRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Validation failed")
	}

	// Get user ID from context (set by auth middleware)
	userID := c.Locals("userID").(uint)

	// Create manga
	manga, err := h.mangaService.CreateManga(&req, userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to create manga")
	}

	return response.Created(c, manga, "Manga created successfully")
}

// GetManga handles GET /api/v1/mangas/:id
func (h *MangaHandler) GetManga(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid manga ID")
	}

	manga, err := h.mangaService.GetMangaByID(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err, "Manga not found")
	}

	return response.Success(c, manga, "Manga retrieved successfully")
}

// GetMangas handles GET /api/v1/mangas
func (h *MangaHandler) GetMangas(c *fiber.Ctx) error {
	mangas, err := h.mangaService.GetMangas()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get mangas")
	}

	return response.Success(c, mangas, "Mangas retrieved successfully")
}

// GetMangasByUser handles GET /api/v1/mangas/user/:userID
func (h *MangaHandler) GetMangasByUser(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("userID"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid user ID")
	}

	mangas, err := h.mangaService.GetMangasByUser(uint(userID))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get user mangas")
	}

	return response.Success(c, mangas, "User mangas retrieved successfully")
}

// UpdateManga handles PUT /api/v1/mangas/:id
func (h *MangaHandler) UpdateManga(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid manga ID")
	}

	var req domain.UpdateMangaRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid request body")
	}

	// Validate request
	if err := validator.ValidateStruct(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Validation failed")
	}

	// Get user ID from context (set by auth middleware)
	userID := c.Locals("userID").(uint)

	// Update manga
	manga, err := h.mangaService.UpdateManga(uint(id), &req, userID)
	if err != nil {
		return response.Error(c, fiber.StatusForbidden, err, "Failed to update manga")
	}

	return response.Success(c, manga, "Manga updated successfully")
}

// DeleteManga handles DELETE /api/v1/mangas/:id
func (h *MangaHandler) DeleteManga(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid manga ID")
	}

	// Get user ID from context (set by auth middleware)
	userID := c.Locals("userID").(uint)

	// Delete manga
	if err := h.mangaService.DeleteManga(uint(id), userID); err != nil {
		return response.Error(c, fiber.StatusForbidden, err, "Failed to delete manga")
	}

	return response.Success(c, map[string]string{"message": "Manga deleted successfully"}, "Manga deleted successfully")
}

// GetActiveMangas handles GET /api/v1/mangas/active
func (h *MangaHandler) GetActiveMangas(c *fiber.Ctx) error {
	mangas, err := h.mangaService.GetActiveMangas()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get active mangas")
	}

	return response.Success(c, mangas, "Active mangas retrieved successfully")
}

// GetMangasByPriceRange handles GET /api/v1/mangas/price?min=0&max=1000
func (h *MangaHandler) GetMangasByPriceRange(c *fiber.Ctx) error {
	minStr := c.Query("min", "0")
	maxStr := c.Query("max", "999999")

	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid min price")
	}

	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid max price")
	}

	mangas, err := h.mangaService.GetMangasByPriceRange(min, max)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get mangas by price range")
	}

	return response.Success(c, mangas, "Mangas by price range retrieved successfully")
}

// GetMangasPaginated handles GET /api/v1/mangas/paginated?page=1&page_size=10
func (h *MangaHandler) GetMangasPaginated(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	// Create pagination request
	pagination := domain.NewPaginationRequest(page, pageSize)

	// Validate pagination
	if err := validator.ValidateStruct(pagination); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid pagination parameters")
	}

	// Get paginated mangas
	result, err := h.mangaService.GetMangasPaginated(pagination)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get paginated mangas")
	}

	return response.Success(c, result, "Paginated mangas retrieved successfully")
}

// GetActiveMangasPaginated handles GET /api/v1/mangas/active/paginated?page=1&page_size=10
func (h *MangaHandler) GetActiveMangasPaginated(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	// Create pagination request
	pagination := domain.NewPaginationRequest(page, pageSize)

	// Validate pagination
	if err := validator.ValidateStruct(pagination); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid pagination parameters")
	}

	// Get paginated active mangas
	result, err := h.mangaService.GetActiveMangasPaginated(pagination)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get paginated active mangas")
	}

	return response.Success(c, result, "Paginated active mangas retrieved successfully")
}

// GetMangasByUserPaginated handles GET /api/v1/mangas/user/:userID/paginated?page=1&page_size=10
func (h *MangaHandler) GetMangasByUserPaginated(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("userID"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid user ID")
	}

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	// Create pagination request
	pagination := domain.NewPaginationRequest(page, pageSize)

	// Validate pagination
	if err := validator.ValidateStruct(pagination); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid pagination parameters")
	}

	// Get paginated user mangas
	result, err := h.mangaService.GetMangasByUserPaginated(uint(userID), pagination)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get paginated user mangas")
	}

	return response.Success(c, result, "Paginated user mangas retrieved successfully")
}

// GetMangasByPriceRangePaginated handles GET /api/v1/mangas/price/paginated?min=0&max=1000&page=1&page_size=10
func (h *MangaHandler) GetMangasByPriceRangePaginated(c *fiber.Ctx) error {
	minStr := c.Query("min", "0")
	maxStr := c.Query("max", "999999")

	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid min price")
	}

	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid max price")
	}

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	// Create pagination request
	pagination := domain.NewPaginationRequest(page, pageSize)

	// Validate pagination
	if err := validator.ValidateStruct(pagination); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err, "Invalid pagination parameters")
	}

	// Get paginated mangas by price range
	result, err := h.mangaService.GetMangasByPriceRangePaginated(min, max, pagination)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err, "Failed to get paginated mangas by price range")
	}

	return response.Success(c, result, "Paginated mangas by price range retrieved successfully")
}
