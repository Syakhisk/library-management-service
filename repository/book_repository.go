package repository

import (
	"library-management/shared/dto"

	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	GetBooks() ([]dto.Book, error)
	CreateBook(book dto.Book) (dto.Book, error)
}

type bookRepository struct {
	db *sqlx.DB
}

func (b *bookRepository) GetBooks() ([]dto.Book, error) {
	// books := []dto.Book{}
	books := make([]dto.Book, 0)
	err := b.db.Select(&books, "SELECT * FROM books")

	return books, err
}

func (b *bookRepository) CreateBook(book dto.Book) (dto.Book, error) {
	var id int
	err := b.db.QueryRow("INSERT INTO books (title, description) VALUES ($1, $2) RETURNING id", book.Title, book.Description).Scan(&id)
	if err != nil {
		return book, err
	}

	return dto.Book{
		ID:          id,
		Title:       book.Title,
		Description: book.Description,
	}, nil
}

func NewBookRepository(db *sqlx.DB) BookRepository {
	// return nya harus pointer,
	// karena butuh instance nya instead of struct nya
	return &bookRepository{
		db,
	}
}
