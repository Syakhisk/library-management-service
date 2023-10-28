package handler

import (
	"library-management/service"
	"library-management/shared"
	"library-management/shared/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	service service.BookService
}

func (b *BookHandler) GetBooks(ctx *fiber.Ctx) error {
	books, err := b.service.GetBooks()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(shared.FormatResponse(nil, err))
	}

	return ctx.Status(http.StatusOK).JSON(shared.FormatResponse(books, nil))
}

func (b *BookHandler) CreateBook(ctx *fiber.Ctx) error {
	var body dto.CreateBookRequestDTO

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(shared.FormatResponse(nil, shared.ErrEmptyBody))
	}

	err = shared.Validate(&body)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(shared.FormatResponse(nil, err))
	}

	book, err := b.service.CreateBook(body)
	if err != nil {
		// return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(shared.FormatResponse(nil, shared.ErrCreateBook))
	}

	return ctx.Status(http.StatusOK).JSON(shared.FormatResponse(book, nil))
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}
