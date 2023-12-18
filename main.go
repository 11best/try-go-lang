package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Book struct {
	ID     int    `json:"id"` //`json:id`(metadata) its for fiber to map request and response to struct
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "Papuan 101", Author: "Papuan"})
	books = append(books, Book{ID: 2, Title: "Patoo", Author: "Papuan"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
	app.Get("/hello", sayHello)

	app.Get("/config", getEnv)

	app.Listen(":8080")
}

func sayHello(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Name": "best",
	})
}

func getEnv(c *fiber.Ctx) error {
	if value, exist := os.LookupEnv("SECRET"); exist {
		return c.JSON(fiber.Map{
			"SECRET": value,
		})
	}

	return c.JSON(fiber.Map{
		"SECRET": "default",
	})
}
