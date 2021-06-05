package main

import "fmt"

type SliceInt []int //phải định nghĩa một kiểu custom thì mới tạo được receiver

/* Receiver không cho phép
func (p []int) DoubleIt() {
	for i, v := range p {
		p[i] = v * 2
	}
}
*/
func (p SliceInt) DoubleIt() {
	for i, v := range p {
		p[i] = v * 2
	}
}

func DemoSliceIntReceiver() {
	var sliceInt SliceInt = []int{1, 2, 3, 4}
	//sliceInt := []int{1, 2, 3, 4}
	sliceInt.DoubleIt()
	fmt.Println(sliceInt)
}
