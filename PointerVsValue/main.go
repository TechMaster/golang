package main

import "fmt"

type Account struct {
	id    string
	name  string
	email string
	pass  string
}

func main() {
	acc := Account{
		id:    "OX13",
		name:  "Cuong",
		email: "cuong@gmail.com",
		pass:  "abc123",
	}

	PassStructAsValue(acc)
	fmt.Println(acc.name) //Cuong

	PassStructAsPointer(&acc)
	fmt.Println(acc.name) //John

}

func PassStructAsValue(acc Account) {
	acc.name = "John"
}

func PassStructAsPointer(acc *Account) {
	acc.name = "John"
}
