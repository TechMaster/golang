package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	ArgumentToPrintlnEscapeToHeap()

	WillSliceEscapeToHeap()

	WillReceiverEscapeToHeap()

	SmallSliceToStack()

	LargeSliceToHeap()

	result := ReturnASlice(10)
	result = result[:0]

	person := ReturnAStructFromFunc()
	person.increaseAge()

	person2 := ReturnPointerStructFromFunc()
	person2.increaseAge()

	InterfaceArgumentEscapeToHeap()

	seasons := ReturnArrayFromFunc()
	fmt.Println(len(seasons))
}
