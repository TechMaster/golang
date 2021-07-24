	package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) String() string {
	return fmt.Sprintf("%v is %v years old", p.FullName(), p.Age)
}

func NewPerson(firstName string, lastName string, age int) *Person {
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
