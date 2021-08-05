package main

import (
	"demofiber/controller"
	"demofiber/errors"
	"demofiber/logger"

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

	app.Get("/", homepage)

	app.Get("/error", genericError) // Trả về lỗi trên màn hình

	app.Get("/login", demoUnAuthorized) //Trả về lỗi JSON

	app.Get("/apierror", demoRESTAPIError) //Trả về lỗi JSON

	app.Get("/dividezero", dividezero) //Trả về custom Error page

	app.Get("/posts", getPosts) //Trả về custom Error page

	app.Get("/page", demoPage) //Trả về custom Error page

	//Để hàm này dưới cùng để bắt lỗi 404 Not Found
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Đường dẫn không tìm thấy")
	})
	app.Listen(":3000")
}

func homepage(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func genericError(c *fiber.Ctx) error {
	return errors.New("Generic Error")
}

func demoUnAuthorized(c *fiber.Ctx) error {
	return errors.UnAuthorized("Không thể xác định danh tính")
}
func demoRESTAPIError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON("Bad Request")
}

func dividezero(c *fiber.Ctx) error {
	return errors.BadRequest("Chia cho 0")
}

func demoPage(c *fiber.Ctx) error {
	viewData := make(fiber.Map)
	viewData["Title"] = "Demo"
	return c.Render("khoahoc/index", viewData)
}

func getPosts(c *fiber.Ctx) error {
	if posts, err := controller.GetAllPosts(); err != nil {
		return err
	} else {
		return c.Render("post", fiber.Map{
			"Posts": posts,
		})
	}

}

// Chuyên xử lý các err mà controller trả về
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	var statusCode = 500

	if e, ok := err.(*fiber.Error); ok { //Thử kiểm tra xem có phải là kiểu fiber.Error không
		statusCode = e.Code
	} else if e, ok := err.(*errors.Error); ok { //Thử kiểm tra xem có phải là kiểu errors.Error
		statusCode = e.Code
	}

	logger.Log.Error(err)

	if err = ctx.Render("error", fiber.Map{
		"ErrorMessage": err.Error(),
		"StatusCode":   statusCode,
	}); err != nil {
		return ctx.Status(500).SendString("Internal Server Error")
	}

	return nil
}
