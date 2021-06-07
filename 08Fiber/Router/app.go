package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//Ví dụ này để hiểu thêm về fiber.Router interface
	router1 := app.Get("/chain", func(c *fiber.Ctx) error {
		fmt.Println("It is first handler")
		return c.Next()
	})

	router2 := router1.Get("/chain", func(c *fiber.Ctx) error {
		fmt.Println("It is second handler")
		return c.Next()
	})

	router2.Get("/chain", func(c *fiber.Ctx) error {
		fmt.Println("It is third handler")
		return c.SendString("Fiber Router is the interface. Any struct can implement it then return itself")
	})

	stacks := app.Stack()
	for _, stack := range stacks {
		for _, route := range stack {
			fmt.Println(route.Method, " : ", route.Path)

			for _, handler := range route.Handlers {
				fmt.Println("   ", handler)
			}
		}
	}

	app.Listen(":3000")
}
