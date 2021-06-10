package main

import (
	"github.com/TechMaster/golang/08Fiber/Repository/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	bookRouter := app.Group("/api/book")
	routes.ConfigBookRouter(&bookRouter) //http://localhost:3000/api/book

	app.Listen(":3000")
}
