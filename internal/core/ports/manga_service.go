package ports

import "github.com/thitiphongD/my-backend/internal/core/domain"

// MangaService defines the interface for manga business operations
type MangaService interface {
	CreateManga(req *domain.CreateMangaRequest, userID uint) (*domain.Manga, error)
	GetMangaByID(id uint) (*domain.Manga, error)
	GetMangas() ([]*domain.Manga, error)
	GetMangasByUser(userID uint) ([]*domain.Manga, error)
	UpdateManga(id uint, req *domain.UpdateMangaRequest, userID uint) (*domain.Manga, error)
	DeleteManga(id uint, userID uint) error
	GetActiveMangas() ([]*domain.Manga, error)
	GetMangasByPriceRange(min, max float64) ([]*domain.Manga, error)

	// Paginated operations
	GetMangasPaginated(pagination *domain.PaginationRequest) (*domain.PaginatedResult[*domain.Manga], error)
	GetActiveMangasPaginated(pagination *domain.PaginationRequest) (*domain.PaginatedResult[*domain.Manga], error)
	GetMangasByUserPaginated(userID uint, pagination *domain.PaginationRequest) (*domain.PaginatedResult[*domain.Manga], error)
	GetMangasByPriceRangePaginated(min, max float64, pagination *domain.PaginationRequest) (*domain.PaginatedResult[*domain.Manga], error)
}
