package router

import (
	"demofiber/controller"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", controller.Homepage)
	routeDemoErrorHandling(app)
	routeSession(app)
	demoParamQuery(app)
}
func routeDemoErrorHandling(app *fiber.App) {
	app.Get("/err/error", controller.GenericError) // Trả về lỗi trên màn hình

	app.Get("/err/login", controller.DemoUnAuthorized) //Trả về lỗi JSON

	app.Get("/err/apierror", controller.DemoRESTAPIError) //Trả về lỗi JSON

	app.Get("/err/loginapi", controller.LoginAPI) //Trả về lỗi JSON

	app.Get("/err/wrap", controller.LogOutAPI) //Trả về lỗi JSON

	app.Get("/err/dividezero", controller.DivideZero) //Trả về custom Error page

	app.Get("/err/posts", controller.GetAllPosts) //Trả về custom Error page

	app.Get("/err/page", controller.DemoPage) //Trả về custom Error page

	app.Get("/err/panic", controller.DemoPanicError) //Xử lý lỗi panic error

	app.Get("/err/warn", controller.DemoWarning) //Xử lý Warning

	app.Get("/err/cause", controller.DemoCause)
}

func demoParamQuery(app *fiber.App) {
	app.Get("/para/author/:name/book/:id", func(c *fiber.Ctx) error {
		fmt.Println(c.Route().Path)

		if id, err := c.ParamsInt("id"); err == nil {
			fmt.Fprintf(c, "%d\n", id)
			return c.SendString(fmt.Sprintf("%s - %d", c.Params("name"), id))
		} else {
			return c.SendString(fmt.Sprintf("%s - %s", c.Params("name"), c.Params("id")))
		}
	})

	// Plus - greedy - not optional
	app.Get("/para/user/+", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("+"))
	})

	// Optional parameter
	app.Get("/para/foo/:name?", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name"))
	})

	// Wildcard - greedy - optional
	app.Get("/para/bar/*", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("*"))
	})

	// This route path will match requests to "/para/name:customVerb", since the parameter character is escaped
	app.Get("/para/name\\:customVerb", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Community")
	})
}

func routeSession(app *fiber.App) {
	app.Get("/session/getsessionid", controller.GetSessionID)
	app.Get("/session/counter", controller.SessionCounter)
	app.Get("/session/authen", controller.SessionAuthenticate)
}
