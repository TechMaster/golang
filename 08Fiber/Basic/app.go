package main

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
) //import thư viện fiber version 2

func main() {
	app := fiber.New()

	app.Get("/", hello)

	app.Get("/:name", sayName) // GET /john

	app.Get("/bye/:name", log, bye)

	app.Listen(":3000")
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func sayName(c *fiber.Ctx) error {
	name, err := url.PathUnescape(c.Params("name"))
	fmt.Println(err)
	fmt.Println(name)

	msg := fmt.Sprintf("Hello, %s 👋!", name)
	return c.SendString(msg) // => Hello john 👋!
}

func log(c *fiber.Ctx) error {
	fmt.Println("Log: " + c.Params("name"))
	return c.Next()
}

func bye(c *fiber.Ctx) error {
	msg := fmt.Sprintf("good bye %s 👋!", c.Params("name"))
	return c.SendString(msg) // => good bye john 👋!
}
