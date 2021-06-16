package controller

import (
	"fmt"

	"github.com/TechMaster/golang/11FiberCRUD/Repository/model"
	repo "github.com/TechMaster/golang/11FiberCRUD/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllBook(c *fiber.Ctx) error {
	return c.JSON(repo.BookRepo.GetAllBooks())
}

func GetBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	book, err := repo.BookRepo.FindBookById(int64(id))
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
	err = repo.BookRepo.DeleteBookById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete book successfully")
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

	bookId := repo.BookRepo.CreateNewBook(book)
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

	err = repo.BookRepo.UpdateBook(updatedBook)
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

	id := repo.BookRepo.Upsert(book)
	return c.SendString(fmt.Sprintf("Book with id = %d is successfully upserted", id))
}
