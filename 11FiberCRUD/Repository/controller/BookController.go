package controller

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
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

func DeleteBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Books.DeleteBookById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateBook(c *fiber.Ctx) error {
	book := new(model.Book)

	err := c.BodyParser(&book)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	bookId := repo.Books.CreateNewBook(book)
	return c.SendString(fmt.Sprintf("New book is created successfully with id = %d", bookId))

}

func UpdateBook(c *fiber.Ctx) error {
	updatedBook := new(model.Book)

	err := c.BodyParser(&updatedBook)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Books.UpdateBook(updatedBook)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Book with id = %d is successfully updated", updatedBook.Id))

}

func UpsertBook(c *fiber.Ctx) error {
	book := new(model.Book)

	err := c.BodyParser(&book)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Books.Upsert(book)
	return c.SendString(fmt.Sprintf("Book with id = %d is successfully upserted", id))
}
