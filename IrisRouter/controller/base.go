package controller

import "github.com/kataras/iris/v12"

func ShowHomePage(ctx iris.Context) {
	_ = ctx.View("index")
}
