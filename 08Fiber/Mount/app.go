package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Title  string  `json:"name"`
	Author string  `json:"author"`
	Rating float32 `json:"rating"`
}

type Book2 struct {
	Title   string   `json:"name"`
	Authors []string `json:"authors"`
	Rating  float32  `json:"rating"`
}

func buildRestV1() (appREST *fiber.App) {
	appREST = fiber.New()
	appREST.Get("/books", func(c *fiber.Ctx) error {
		var books []Book
		books = append(books, Book{Title: "Dế Mèn Phiêu Lưu Ký", Author: "Tô Hoài", Rating: 4.5})
		books = append(books, Book{Title: "Clean Code", Author: "Robert Cecil Martin", Rating: 4.4})
		return c.JSON(books)
	})
	return
}

func buildRestV2() (appREST *fiber.App) {
	appREST = fiber.New()
	appREST.Get("/books", func(c *fiber.Ctx) error {
		var books []Book2
		books = append(books, Book2{Title: "Dế Mèn Phiêu Lưu Ký", Authors: []string{"Tố Hoài", "Mark Rush"}, Rating: 4.5})
		books = append(books, Book2{Title: "Clean Code", Authors: []string{"Robert Cecil Martin", "David James"}, Rating: 4.4})
		return c.JSON(books)
	})

	return
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})
	app.Mount("/v1", buildRestV1()) //http://localhost:3000/v1/books
	app.Mount("/v2", buildRestV2()) //http://localhost:3000/v2/books

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})
	app.Listen(":3000")
}
