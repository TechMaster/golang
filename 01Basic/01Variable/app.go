package main

import (
	"fmt"
	"strconv"

	"github.com/TechMaster/golang/01Basic/01Variable/util"
)

//Khai báo biến toàn cục (global variable)
var company_name = "techmaster"
var current_year int

func main() {
	var name = "John"
	var lastname = "Smith"
	age := 30

	fmt.Println(name + " " + lastname + " is " + strconv.Itoa(age))

	s := fmt.Sprintf("%s %s is %d\n", name, lastname, age)
	fmt.Printf(s)
	fmt.Printf("%s %s is %d\n", name, lastname, age)

	current_year = 2021
	fmt.Printf("%s %d\n", company_name, current_year)

	fmt.Println(demoString)
	fmt.Println(util.UtilString)
	//fmt.Println(util.utilString)
}
