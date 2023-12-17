package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"` //`json:id`(metadata) its for fiber to map request and response to struct
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Papuan 101", Author: "Papuan"})
	books = append(books, Book{ID: 2, Title: "Patoo", Author: "Papuan"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Listen(":8080")
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}
