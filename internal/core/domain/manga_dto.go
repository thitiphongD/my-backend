package domain

// CreateMangaRequest represents the request body for creating a manga
type CreateMangaRequest struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required,min=0"`
	IsActive bool    `json:"is_active"`
}

// UpdateMangaRequest represents the request body for updating a manga
type UpdateMangaRequest struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required,min=0"`
	IsActive bool    `json:"is_active"`
}

// MangaResponse represents manga data for API responses
type MangaResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	IsActive    bool    `json:"is_active"`
	UserCreated uint    `json:"user_created"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
