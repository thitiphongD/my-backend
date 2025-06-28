package domain

// PaginationRequest represents pagination parameters from request
type PaginationRequest struct {
	Page     int `query:"page" validate:"min=1"`
	PageSize int `query:"page_size" validate:"min=1,max=100"`
}

// PaginationResponse represents pagination metadata in response
type PaginationResponse struct {
	CurrentPage  int   `json:"current_page"`
	PageSize     int   `json:"page_size"`
	TotalItems   int64 `json:"total_items"`
	TotalPages   int   `json:"total_pages"`
	HasNextPage  bool  `json:"has_next_page"`
	HasPrevPage  bool  `json:"has_prev_page"`
	NextPage     *int  `json:"next_page,omitempty"`
	PreviousPage *int  `json:"previous_page,omitempty"`
}

// PaginatedResult represents a paginated result with data and metadata
type PaginatedResult[T any] struct {
	Data       []T                 `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

// NewPaginationRequest creates a new pagination request with default values
func NewPaginationRequest(page, pageSize int) *PaginationRequest {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10 // Default page size
	}
	return &PaginationRequest{
		Page:     page,
		PageSize: pageSize,
	}
}

// GetOffset calculates the offset for database queries
func (p *PaginationRequest) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// GetLimit returns the page size as limit
func (p *PaginationRequest) GetLimit() int {
	return p.PageSize
}

// NewPaginationResponse creates pagination metadata
func NewPaginationResponse(page, pageSize int, totalItems int64) *PaginationResponse {
	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	hasNextPage := page < totalPages
	hasPrevPage := page > 1

	var nextPage *int
	var prevPage *int

	if hasNextPage {
		next := page + 1
		nextPage = &next
	}

	if hasPrevPage {
		prev := page - 1
		prevPage = &prev
	}

	return &PaginationResponse{
		CurrentPage:  page,
		PageSize:     pageSize,
		TotalItems:   totalItems,
		TotalPages:   totalPages,
		HasNextPage:  hasNextPage,
		HasPrevPage:  hasPrevPage,
		NextPage:     nextPage,
		PreviousPage: prevPage,
	}
}
