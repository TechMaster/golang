package routerV3

import (
	"github.com/TechMaster/golang/08Fiber/01Basic/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigRouterV3(router fiber.Router) {

	router.Get("/list", controller.ListV3)

	router.Get("/show", func(c *fiber.Ctx) error {
		return c.SendString("about")
	})
}
