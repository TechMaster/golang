package main

import "fmt"

type Person struct {
	FirstName string //kiểu đứng sau tên biến
	LastName  string
	Age       int
	Address1  Address  //value
	Address2  *Address //pointer
}

type Address struct {
	Country string
	City    string
}

//Truyền vào con trỏ tham chiếu
func (p *Person) FullName() string {
	fmt.Printf("pointer receiver %p\n", p) //In ra địa chỉ trong vùng nhớ của con trỏ p
	p.Age = 100
	return p.FirstName + " " + p.LastName
}

//Truyền vào biến
func (p Person) String() string {
	fmt.Printf("value receiver : %p\n", &p) //In ra địa chỉ trong vùng nhớ của biến p
	p.Age = 200
	return fmt.Sprintf("%v is %v years old", p.FirstName, p.Age)
}

func NewAPerson(firstName string, lastName string, age int) *Person {
	if age < 0 {
		return nil
	}
	p := new(Person)
	p.FirstName = firstName
	p.LastName = lastName
	p.Age = age
	return p
}

//-----Fluent API
func BuildPerson() *Person {
	return new(Person)
}

func (p *Person) WithFirstName(firstName string) *Person {
	p.FirstName = firstName
	return p
}

func (p *Person) WithLastName(lastName string) *Person {
	p.LastName = lastName
	return p
}

func (p *Person) WithAge(age int) *Person {
	p.Age = age
	return p
}
