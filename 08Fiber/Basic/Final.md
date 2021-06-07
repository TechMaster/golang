Mã nguồn cuối cùng

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", hello)

	app.Get("/:name", sayName) // GET /john

	app.Get("/bye/:name", log, bye)

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
		return c.SendString("Fiber Router is the interface. Any struct can implement it then ret ")
	})

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
```