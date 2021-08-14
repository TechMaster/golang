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

	data := map[string]interface{}{
		"name":  "rock",
		"email": "rock@gmail.com",
		"pass":  "abc123",
	}
	PassMapAsValue(data)
	fmt.Println(data["name"]) //đã đổi thành John

	PassMapAsPointer(&data)
	fmt.Println(data["name"]) //đã đổi thành Hann
}

func PassStructAsValue(acc Account) {
	acc.name = "John"
}

func PassStructAsPointer(acc *Account) {
	acc.name = "John"
}

func PassMapAsValue(data map[string]interface{}) {
	data["name"] = "John"
	data["email"] = "john@gmail.com"
}

func PassMapAsPointer(data *map[string]interface{}) {
	(*data)["name"] = "Hann"
	(*data)["email"] = "hann@gmail.com"
}
