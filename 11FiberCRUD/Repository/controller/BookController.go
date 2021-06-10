package controller

import (
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllBook(c *fiber.Ctx) error {
	return c.JSON(repo.Books.GetAllBooks())
}

func GetBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	book, err := repo.Books.FindBookById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(book)
}
