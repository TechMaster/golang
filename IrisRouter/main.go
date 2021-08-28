package main

import (
	"fmt"
	"irouter/router"
	"irouter/template"

	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	logFile := logger.Init() //Cần phải có 2 file error.html và info.html ở /views
	if logFile != nil {
		defer logFile.Close()
	}

	app.WrapRouter(routerWrapper)
	app.UseRouter(routerMiddleware)

	app.UseGlobal(globalMiddleware) //Tiến hành kiểm tra phân quyền ở đây
	app.Use(useMiddleware)
	app.UseError(errorMiddleware)
	// app.Done(done)
	// app.DoneGlobal(doneGlobal)

	// Adding a OnErrorCode(iris.StatusNotFound) causes `.UseGlobal`
	// to be fired on 404 pages without this,
	// only `UseError` will be called, and thus should
	// be used for error pages.
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	router.RegisterRoute(app)
	template.InitViewEngine(app)
	for _, route := range app.GetRoutes() {
		fmt.Println(route.Name)
	}
	_ = app.Listen(":3000")
}
