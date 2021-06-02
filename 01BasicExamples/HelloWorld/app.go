package main

import (
	"fmt"

	"github.com/techmaster.vn/helloworld/utils"
)

func say(msg string) {
	fmt.Println(msg)
}
func main() {
	say("Hello World")
	utils.Say("Hi World")
	utils.Foo("Hi World")
}
