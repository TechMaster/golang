package controller

import "github.com/gofiber/fiber/v2"

func ListV3(c *fiber.Ctx) error {
	return c.SendString("list")
}
