package router

import (
	"demofiber/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", controller.Homepage)

	app.Get("/error", controller.GenericError) // Trả về lỗi trên màn hình

	app.Get("/login", controller.DemoUnAuthorized) //Trả về lỗi JSON

	app.Get("/apierror", controller.DemoRESTAPIError) //Trả về lỗi JSON

	app.Get("/dividezero", controller.DivideZero) //Trả về custom Error page

	app.Get("/posts", controller.GetAllPosts) //Trả về custom Error page

	app.Get("/page", controller.DemoPage) //Trả về custom Error page

	app.Get("/panic", controller.DemoPanicError) //Xử lý lỗi panic error

	app.Get("/warn", controller.DemoWarning) //Xử lý Warning
	app.Get("/cause", controller.DemoCause)

	app.Get("/getsessionid", controller.GetSessionID)

	app.Get("/counter", controller.SessionCounter)
}
