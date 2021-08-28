package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/TechMaster/core/template"
	"github.com/kataras/iris/v12"
)

func Benchmark_HTML(b *testing.B) {
	/*engine := iris.HTML("./views", ".html")
	 */

	app := iris.New() // defaults to these
	template.InitHTMLEngine(app, "./views", "")
	engine := template.HTMLEngine
	engine.Load()

	data := initData()
	buf := new(bytes.Buffer)
	if err := engine.ExecuteWriter(buf, "index", "", data); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(buf.String())
}

func Benchmark_Block(b *testing.B) {
	/*engine := blocks.NewBlocks("./views", ".html")
	_ = engine.Load()*/

	app := iris.New() // defaults to these
	template.InitBlockEngine(app, "./views", "")
	engine := template.BlockEngine

	data := initData()
	buf := new(bytes.Buffer)
	if err := engine.ExecuteWriter(buf, "index", "", data); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(buf.String())
}
