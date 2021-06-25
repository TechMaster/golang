package controller

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
	"github.com/TechMaster/golang/15GoMySQL/repo"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []model.Product

	result := repo.Db.Preload("Country").
		Preload("Manufacturer").
		Preload("Category").
		Find(&products)

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(products)
	}
}
