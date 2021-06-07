package main

import (
	"github.com/gofiber/fiber/v2"
)

var repository Repository

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
		return c.JSON(repository.GetAllBooks())
	})

	appREST.Get("/books/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		book, err := repository.FindBookById(id)
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.JSON(book)
	})

	return
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
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

	repository = Repository{}
	repository.InitData("sql:534444")

	app.Listen(":3000")
}
