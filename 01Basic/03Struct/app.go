package main

import "fmt"

func main() {
	person := Person{"Trinh", "Cuong", 45,
		Address{"Vietnam", "Hanoi"},
		&Address{"USA", "California"}}

	fmt.Printf("Địa chỉ con trỏ ban đầu %p\n", &person)

	fmt.Println(person.FullName())

	//fmt.Println(person.Age)

	fmt.Println(person.String())
	//fmt.Println(person.Age)

	//Sử dụng constructor. Trong Golang có cú pháp constructor chuẩn
	//mà lập trình viên tự định nghĩa constructor
	//Chúng ta có thể đặt tên hàm constructor là NewPerson hay NewAPerson
	//Ngược lại trong Java tên constructor luôn phải trùng tên với Java constructor
	/*	tom := NewAPerson("Trinh", "Cuong", -2)
		if tom != nil {
			fmt.Println(tom.FullName())
		}

		//Fluent API bản chất chỉ là cách viết nối tiếp các method mà thôi
		jack := BuildPerson().WithFirstName("Jack").WithLastName("London").WithAge(12)
		fmt.Println(jack)*/
}
