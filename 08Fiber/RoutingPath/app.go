package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// http://localhost:3000/hello/cuong.trinh
	// Hello, cuong.trinh 👋!
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	// http://localhost:3000/person/John/15
	// 👴 John is 15 years old
	app.Get("/person/:name/:age", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// http://localhost:3000/file/dictionary.txt
	// 📃 dictionary.txt
	app.Get("/file/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// http://localhost:3000/flights/LAX-SFO
	// 💸 From: LAX, To: SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET http://localhost:3000/api/register
	//✋ register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	//http://localhost:3000/user/cuong/books/dictionary
	//user: cuong, book: dictionary

	app.Get("/user/:name/books/:title", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("user: %s, book: %s", c.Params("name"), c.Params("title"))
		return c.SendString(msg)
	})

	//http://localhost:3000/shoes?order=desc&brand=nike
	app.Get("/shoes", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("order: %s - brand: %s", c.Query("order"), c.Query("brand"))
		return c.SendString(msg)
	})

	stacks := app.Stack()
	for _, stack := range stacks {
		for _, route := range stack {
			fmt.Println(route.Method, " : ", route.Path)
		}
	}

	log.Fatal(app.Listen(":3000"))
}
