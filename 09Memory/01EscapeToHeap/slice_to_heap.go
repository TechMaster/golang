package main

import "fmt"

func WillSliceEscapeToHeap() {
	intArr := [6]int{1, 2, 3, 4, 5, 6}
	for _, i := range intArr {
		fmt.Println(i) //i escapes to heap
	}
	fmt.Println(intArr) //intArr escapes to heap
}
