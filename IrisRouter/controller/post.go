package controller

import (
	"fmt"

	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)

func GetAllPost(ctx iris.Context) {
	_, _ = ctx.WriteString("Get All Posts")
}

func GetPostById(ctx iris.Context) {
	if id, err := ctx.Params().GetInt("id"); err != nil {
		logger.Log(ctx, eris.Warning("Cannot find post"))
	} else {
		_, _ = ctx.WriteString(fmt.Sprintf("Get Post.id = %d", id))
	}
}
func CreatePost(ctx iris.Context) {
	_, _ = ctx.WriteString("Creat Post")
}
