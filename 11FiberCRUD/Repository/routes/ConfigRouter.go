package routes

import (
	"github.com/TechMaster/golang/08Fiber/Repository/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigBookRouter(router *fiber.Router) {
	//Return all books
	(*router).Get("/", controller.GetAllBook)

	(*router).Get("/:id", controller.GetBookById)

	(*router).Delete("/:id", controller.DeleteBookById)

	(*router).Post("", controller.CreateBook)
}
