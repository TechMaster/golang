package main

import "fmt"

func (p *Person) increaseAge() { //p does not escape
	p.Age += 1
}

func (p *Person) reverseName() (result string) { //p does not escape
	for _, s := range p.Name {
		result = string(s) + result //string(s) + result escapes to heap but string(s) does not escape
	}
	return
}

func makeAPerson(name string, age int) (person *Person) {
	person = new(Person) //new(Person) escapes to heap
	person.Name = name
	person.Age = age
	return
}

func WillReceiverEscapeToHeap() {
	alice := new(Person) //new(Person) does not escape
	alice.Name = "Alice"
	alice.Age = 21
	alice.increaseAge()
	temp := alice.reverseName()
	fmt.Println(len(temp)) //len(temp) escapes to heap because it pass as argument in fmt.Println

	bob := makeAPerson("Bob", 22) //new(Person) does not escape
	bob.reverseName()
}
