package main

import (
	"fmt"
	"pointer_value/value"
)

type Account struct {
	id    string
	name  string
	email string
	pass  string
}

func main() {
	demo := map[string]interface{}{}
	if demo["name"] != nil {
		fun := demo["name"].(string)
		fmt.Println(fun)
	}

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
	//------------------------------------------------
	var accNewV = value.AccountNew{
		Email:    "test@example.com",
		FullName: "Test Name",
		Password: "Real Password",
	}
	fmt.Println(len(value.AccountRepository.GetAll()))
	if id, err := value.AccountRepository.Save(accNewV); err == nil {
		fmt.Println(id)
	}
	fmt.Println(len(value.AccountRepository.GetAll()))

}

//Pass by Value
func PassStructAsValue(acc Account) {
	acc.name = "John"
}

//Pass by Pointer
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
