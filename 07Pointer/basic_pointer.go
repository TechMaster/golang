package main

import "fmt"

func DemoBasicPointer() {
	age := new(int) //age là con trỏ
	*age = 26

	fmt.Println(*age)
	fmt.Println(age)
	fmt.Println(&age)

	tuoi := 27
	age = &tuoi //Toán tử lấy địa chỉ
	fmt.Println(*age)
	*age += 1
	fmt.Println(tuoi)
}
