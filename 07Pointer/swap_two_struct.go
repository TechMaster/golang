package main

import (
	"fmt"
)

func SwapPerson(a *(*Person), b *(*Person)) {
	fmt.Printf("a = %p\n", a)
	fmt.Printf("b = %p\n", b)
	a, b = b, a
	fmt.Printf("a = %p\n", a)
	fmt.Printf("b = %p\n", b)
	fmt.Println("alice ", *a)
	fmt.Println("bob ", *b)
}

func DemoSwapTwoStructs() {
	alice := Person{Name: "Alice", Age: 18}
	bob := Person{Name: "Bob", Age: 19}

	alice, bob = bob, alice
	fmt.Println("alice ", alice)
	fmt.Println("bob ", bob)

}
