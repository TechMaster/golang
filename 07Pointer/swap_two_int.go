package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func DemoSwapTwoInts() {
	x, y := 5, 10
	fmt.Println("Before SWAP: ", x, y)
	swap(&x, &y)
	fmt.Println("After SWAP:  ", x, y)
}
