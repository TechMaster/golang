package main

func ReturnAStructFromFunc() Person {
	person := Person{Name: "Cường", Age: 46}
	return person //does not escape to heap
}

func ReturnPointerStructFromFunc() *Person {
	person := Person{Name: "Cường", Age: 46} //moved to heap: person
	return &person
}
