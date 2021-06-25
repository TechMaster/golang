package controller

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
	"github.com/TechMaster/golang/15GoMySQL/repo"
	"github.com/gofiber/fiber/v2"
)

func GetManufacturers(c *fiber.Ctx) error {
	var manufacturers []model.Manufacturer

	//Hãy bật tắt Preload("Country").
	result := repo.Db.Preload("Country").Find(&manufacturers)
	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(manufacturers)
	}
}
