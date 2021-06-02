package main

import (
	"fmt"

	"github.com/TechMaster/foo"
	"github.com/TechMaster/greeting"
	"github.com/TechMaster/mygomodule"
	"github.com/techmaster.vn/app/utils"
)

func init() {
	fmt.Println("hàm này luôn chạy trước tiên")
}
func main() {
	greeting.Say("Hello world")
	fmt.Println(foo.Bar())
	var result = mygomodule.Reverse("Alo")
	fmt.Println(result)
	utils.Demo()
}
