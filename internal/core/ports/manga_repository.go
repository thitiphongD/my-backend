package ports

import "github.com/thitiphongD/my-backend/internal/core/domain"

// MangaRepository defines the interface for manga data access
type MangaRepository interface {
	// Manga CRUD operations
	Create(manga *domain.Manga) error
	GetByID(id uint) (*domain.Manga, error)
	GetByUserID(userID uint) ([]*domain.Manga, error)
	List() ([]*domain.Manga, error)
	Update(manga *domain.Manga) error
	Delete(id uint) error

	// Additional queries
	GetActiveMangas() ([]*domain.Manga, error)
	GetMangasByPriceRange(min, max float64) ([]*domain.Manga, error)

	// Paginated queries
	ListPaginated(pagination *domain.PaginationRequest) ([]*domain.Manga, int64, error)
	GetActiveMangasPaginated(pagination *domain.PaginationRequest) ([]*domain.Manga, int64, error)
	GetMangasByUserIDPaginated(userID uint, pagination *domain.PaginationRequest) ([]*domain.Manga, int64, error)
	GetMangasByPriceRangePaginated(min, max float64, pagination *domain.PaginationRequest) ([]*domain.Manga, int64, error)
}
