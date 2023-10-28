package service

import (
	"database/sql"
	"library-management/repository"
	"library-management/shared"
	"library-management/shared/dto"
)

type BookService interface {
	GetBooks() (dto.GetBookResponseDTO, error)
	CreateBook(data dto.CreateBookRequestDTO) (dto.CreateBookResponseDTO, error)
}

type bookService struct {
	repository repository.BookRepository
}

func (b *bookService) GetBooks() (dto.GetBookResponseDTO, error) {
	books, err := b.repository.GetBooks()
	if err != nil {
		return dto.GetBookResponseDTO{}, err
	}

	// mapping entity to dto
	bookResponses := make([]dto.BookResponse, 0)
	for _, book := range books {
		bookEntry := dto.BookResponse{
			ID:    book.ID,
			Title: book.Title,
		}

		if book.Description.Valid {
			bookEntry.Description = book.Description.String
		}

		bookResponses = append(bookResponses, bookEntry)
	}

	return dto.GetBookResponseDTO{
		Books: bookResponses,
	}, nil
}

func (b *bookService) CreateBook(data dto.CreateBookRequestDTO) (dto.CreateBookResponseDTO, error) {
	book := dto.Book{
		Title: data.Title,
	}

	if data.Description != "" {
		book.Description = sql.NullString{
			String: data.Description,
			Valid:  true,
		}
	}

	newBook, err := b.repository.CreateBook(book)
	if err != nil {
		return dto.CreateBookResponseDTO{}, shared.ErrCreateBook
	}

	return dto.CreateBookResponseDTO{
		ID:          newBook.ID,
		Title:       newBook.Title,
		Description: newBook.Description.String,
	}, nil
}

func NewBookService(repository repository.BookRepository) BookService {
	return &bookService{
		repository: repository,
	}
}
