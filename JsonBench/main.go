package main

import (
	"fmt"
	"jsonbench/demo"
)

func main() {
	jsonString := demo.Goccy_Struct2JSON(demo.APerson)
	aPerson := demo.Goccy_JSON2Struct(jsonString)
	mapInterface := demo.Goccy_JSON2Map(jsonString)
	fmt.Println(aPerson.Name)
	fmt.Println(mapInterface["Name"])
}
