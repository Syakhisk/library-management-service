package dto

import "database/sql"

type (
	Book struct {
		ID          int            `db:"id"`
		Title       string         `db:"title"`
		Description sql.NullString `db:"description"`
	}

	BookResponse struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	GetBookResponseDTO struct {
		Books []BookResponse `json:"books"`
	}

	CreateBookRequestDTO struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description,omitempty" validate:"omitempty"`
	}

	CreateBookResponseDTO struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
)
