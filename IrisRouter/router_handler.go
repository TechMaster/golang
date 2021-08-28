package main

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func mainHandler(ctx iris.Context) {
	_, _ = ctx.WriteString("Main Handler")
}

func notFoundHandler(ctx iris.Context) {
	_, _ = ctx.WriteString("404 Error Handler")
}

func routerWrapper(w http.ResponseWriter, r *http.Request,
	router http.HandlerFunc) {
	_, _ = w.Write([]byte("#1 .WrapRouter" + r.RequestURI + "\n"))

	// Continue by executing the Iris Router and let it do its job.
	router(w, r)
}

func routerMiddleware(ctx iris.Context) {
	_, _ = ctx.WriteString("#2 .UseRouter: " + ctx.GetContentTypeRequested() + "\n")
	ctx.Next()
}

/*
Kiểm tra phân quyền ở đây là đẹp
*/
func globalMiddleware(ctx iris.Context) {
	_, _ = ctx.WriteString("#3 .UseGlobal: " + ctx.GetCurrentRoute().String() + "\n")
	ctx.Next()
}

func useMiddleware(ctx iris.Context) {
	_, _ = ctx.WriteString("#4 .Use: " + ctx.GetCurrentRoute().String() + "\n")
	ctx.Next()
}

func errorMiddleware(ctx iris.Context) {
	_, _ = ctx.WriteString("#3 .UseError\n")
	ctx.Next()
}
