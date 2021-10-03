package main

import (
	"github.com/TechMaster/core/logger"
	"github.com/TechMaster/core/template"
	"github.com/TechMaster/eris"
	"github.com/kataras/iris/v12"
)

func initData() map[string]interface{} {
	var data = make(map[string]interface{})
	type person struct {
		Name    string
		Email   string
		Age     int
		Address string
		Status  bool
	}
	type order struct {
		Item   string
		Price  int
		Amount int
	}
	data["user"] = person{
		Name:    "Trịnh Minh Cường",
		Email:   "cuong@techmaster.vn",
		Age:     46,
		Address: "tầng 12A, Viwaseen Tower, 48 Tố Hữu, Nam Từ Liêm, Hà nội",
		Status:  true,
	}

	data["orders"] = []order{
		{
			Item:   "Nike Shoes",
			Price:  1200000,
			Amount: 1,
		},
		{
			Item:   "iPhoneX",
			Price:  18000000,
			Amount: 2,
		},
		{
			Item:   "Vinfast Fadil",
			Price:  450000000,
			Amount: 1,
		},
	}
	return data

}

func main() {
	app := iris.New() // defaults to these

	logFile := logger.Init() //Cần phải có 2 file error.html và info.html ở /views
	if logFile != nil {
		defer logFile.Close()
	}

	template.InitBlockEngine(app, "./views", "")
	//template.InitHTMLEngine(app, "./views", "")

	app.Get("/", hi)
	app.Get("/err", showerr)

	_ = app.Listen(":8002")
}

func hi(ctx iris.Context) {
	_ = ctx.View("index", initData())
}

func showerr(ctx iris.Context) {
	if err := foo(); err != nil {
		logger.Log(ctx, err)
		return
	}

	logger.Info(ctx, "Chạy ngon không lỗi")
}

func foo() error {
	if err := bar(); err != nil {
		return err
	}
	return nil
}

func bar() error {
	return eris.New("Demo")
}
