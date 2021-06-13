package main

import "fmt"

func DemoChannelString() {
	pipe := make(chan string, 2)
	go func() {
		pipe <- "water"
		pipe <- "water"
		close(pipe)
	}()

	for receiver := range pipe {
		fmt.Println(receiver)
	}
}
