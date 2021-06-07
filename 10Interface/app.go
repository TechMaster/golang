package main

import "fmt"

type Animal interface {
	Move()
	Sound()
}

type Cat struct {
	Name string
}

func (cat *Cat) Move() {
	fmt.Println(cat.Name + " moves")
}

func (cat *Cat) Sound() {
	fmt.Println(cat.Name + " sounds meo meo")
}

type Bird struct {
	Name string
}

func (bird *Bird) Move() {
	fmt.Println(bird.Name + " flies")
}

func (bird *Bird) Sound() {
	fmt.Println(bird.Name + " cuckoo cuckoo")
}

func main() {
	animals := make([]Animal, 2)
	animals[0] = &Bird{Name: "Gà"}
	animals[1] = &Cat{Name: "Mèo"}

	for _, animal := range animals {
		animal.Move()
		animal.Sound()
	}
}
