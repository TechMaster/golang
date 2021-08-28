package main

import "fmt"

func main() {

	roles := map[int]interface{}{
		1: true,
		2: true,
		3: false,
	}

	if r1 := roles[1]; r1.(bool) {
		fmt.Println("Correct")
	}

	if r3 := roles[3]; !r3.(bool) {
		fmt.Println("Correct")
	}

	if r4 := roles[4]; r4.(bool) {
		fmt.Println("Correct")
	}

	roles[5] = false
	roles[6] = true

	if r6 := roles[6]; r6.(bool) {
		fmt.Println("Correct")
	}

}
