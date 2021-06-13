package main

import "fmt"

func DemoChannel() {
	sub := make(chan interface{})

	go func(c <-chan interface{}) {
		for data := range c {
			fmt.Println(data)
		}
	}(sub)

	sub <- "Hey there"
//	sub <- "Hi"
//	sub <- "Goo"
//	sub <- "World"
	close(sub)

}
