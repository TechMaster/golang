package template

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/view"
)

var ViewEngine *view.BlocksEngine

func InitViewEngine(app *iris.Application) {
	ViewEngine = iris.Blocks("./views", ".html")
	app.RegisterView(ViewEngine)
}
