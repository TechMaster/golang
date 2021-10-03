package main

import (
	"github.com/TechMaster/core2/foo"
	"github.com/TechMaster/core2/logger"
	"github.com/TechMaster/core2/template"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()        // defaults to these
	logFile := logger.Init() //Cần phải có 2 file error.html và info.html ở /views
	if logFile != nil {
		defer logFile.Close()
	}
	app.Use(foo.DemoError)
	app.Get("/", home)
	app.Get("/foo", callfoo)
	template.InitViewEngine(app)
	_ = app.Listen(":8002")
}

func home(ctx iris.Context) {
	singleton := logger.GetSingleton()
	_, _ = ctx.WriteString(singleton.Name)
}

func callfoo(ctx iris.Context) {
	foo.DemoError(ctx)
}
