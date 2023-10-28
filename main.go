package main

import (
	"github.com/gofiber/fiber/v2"

	"library-management/handler"
	"library-management/internal"
	"library-management/repository"
	"library-management/service"
)

func main() {
	var fiberConfig = fiber.Config{}

	app := fiber.New(fiberConfig)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	db := internal.NewDB()
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	bookRoute := app.Group("/books")
	bookRoute.Get("/", bookHandler.GetBooks)
	bookRoute.Post("/", bookHandler.CreateBook)

	app.Listen(":3000")
}
