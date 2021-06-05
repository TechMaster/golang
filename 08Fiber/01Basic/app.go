package main

import (
	"fmt"
	"log"

	"github.com/TechMaster/golang/08Fiber/01Basic/routerV3"
	"github.com/gofiber/fiber/v2"
)

func middleware(c *fiber.Ctx) error {
	fmt.Println("Don't mind me!")
	return c.Next()
}

func handler(c *fiber.Ctx) error {
	return c.SendString(c.Path())
}

func setupRoutes(app *fiber.App) {
	// Root API route
	api := app.Group("/api", middleware) // /api

	// API v1 routes
	v1 := api.Group("/v1", middleware) // /api/v1
	v1.Get("/list", handler)           // /api/v1/list
	v1.Get("/user", handler)           // /api/v1/user

	// API v2 routes
	v2 := api.Group("/v2", middleware) // /api/v2
	v2.Get("/list", handler)           //
	v2.Get("/user", handler)           // /api/v2/user

	apiV3 := api.Group("/v3") // /api/v2
	routerV3.ConfigRouterV3(apiV3)

}

func main() {
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
