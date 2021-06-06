package main

func LargeSliceToHeap() {
	largeArray := make([]Person, 10000) //make([]Person, 10000) escapes to heap
	for _, person := range largeArray {
		person.Name = ""
		person.Age = 0
	}
}

func SmallSliceToStack() {
	smallArray := make([]Person, 1000) //make([]Person, 1000) does not escape
	for _, person := range smallArray {
		person.Name = ""
		person.Age = 0
	}
}
