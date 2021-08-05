package main

import (
	"demofiber/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
) //import thư viện fiber version 2

func main() {
	app := fiber.New(
		fiber.Config{
			Views:        html.New("./views", ".html"),
			ErrorHandler: CustomErrorHandler, //Đăng ký hàm xử lý lỗi ở đây
		},
	)

	router.RegisterRoutes(app)
	//Để hàm này dưới cùng để bắt lỗi 404 Not Found
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Đường dẫn không tìm thấy")
	})

	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err.Error())
	}
}
