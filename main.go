package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/11best/try-go-lang/docs"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Book struct {
	ID     int    `json:"id"` //`json:id`(metadata) its for fiber to map request and response to struct
	Title  string `json:"title"`
	Author string `json:"author"`
}

// dummy user
var dummyUser = User{
	Email:    "email@mail.com",
	Password: "1234",
}

var books []Book

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load env error")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "Papuan 101", Author: "Papuan"})
	books = append(books, Book{ID: 2, Title: "Patoo", Author: "Papuan"})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Post("/login", login)

	// jwt middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	// log method and time middleware
	app.Use(checkMiddleware)

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
	app.Get("/hello", sayHello)

	app.Get("/config", getEnv)

	app.Listen(":7777")
}

func sayHello(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Name": "best",
	})
}

func getEnv(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("JWT_SECRET"),
	})
}

func checkMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	user := c.Locals("user").(*jwt.Token) // decrypt jet
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}

	fmt.Printf("Role: %s, URL: %s, Method: %s, Time: %s\n",
		claims["role"], c.OriginalURL(), c.Method(), start)

	return c.Next() // continue
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != dummyUser.Email || user.Password != dummyUser.Password {
		return fiber.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (encrypt)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "login success",
		"token":   t,
	})
}
