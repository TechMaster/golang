package router

import (
	"fmt"
	"irouter/controller"

	"github.com/kataras/iris/v12"
)

func RegisterRoute(app *iris.Application) {
	app.Get("/", controller.ShowHomePage)
	private := app.Party("/private")
	{
		fmt.Println(private.GetRelPath())
		private.Get("/post", controller.GetAllPost)
		private.Get("/post/{id:int}", controller.GetPostById)
		private.Post("/post", controller.CreatePost)
	}

	product := app.Party("/product")
	{
		product.Get("/", controller.GetAllProducts)
		product.Put("/{id:int}/{type:string}", controller.UpdateProductById)
		product.Delete("/{id:int}/rock/{color:string}", controller.DeleteProductById)
	}
}
