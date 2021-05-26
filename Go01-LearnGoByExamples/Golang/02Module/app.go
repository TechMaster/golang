package main

import (
	"fmt"

	"github.com/TechMaster/foo"
	"github.com/TechMaster/greeting"
)

func main() {
	greeting.Say("Hello world")
	fmt.Println(foo.Bar())
}
