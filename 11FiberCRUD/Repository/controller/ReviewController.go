package controller

import (
	"fmt"

	"github.com/TechMaster/golang/11FiberCRUD/Repository/model"
	repo "github.com/TechMaster/golang/11FiberCRUD/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllReviews(c *fiber.Ctx) error {
	return c.JSON(repo.ReviewRepo.GetAllReviews())
}

func GetReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	Review, err := repo.ReviewRepo.FindReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(Review)
}

// Cập nhật lại average rating
func DeleteReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.ReviewRepo.DeleteReviewById(int64(id))

	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		
		return c.SendString("delete review successfully")
	}
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	reviewId := repo.ReviewRepo.CreateNewReview(review)
	//Tính toán lại averating của Review ở đây
	return c.SendString(fmt.Sprintf("New review is created successfully with id = %d", reviewId))

}

func UpdateReview(c *fiber.Ctx) error {
	updatedReview := new(model.Review)

	err := c.BodyParser(&updatedReview)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.ReviewRepo.UpdateReview(updatedReview)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Review with id = %d is successfully updated", updatedReview.Id))
}

func UpsertReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.ReviewRepo.UpsertReview(review)
	return c.SendString(fmt.Sprintf("Review with id = %d is successfully upserted", id))
}
