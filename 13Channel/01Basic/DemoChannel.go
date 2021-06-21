package main

import (
	"fmt"
	"time"
)

func DemoChannel0() {
	pipe := make(chan string)

	go func() {
		for receiver := range pipe {
			fmt.Println(receiver)
		}
	}()

	pipe <- "water 1"
	pipe <- "water 2"
	pipe <- "water 3"
	//close(pipe)
	time.Sleep(time.Millisecond)
}

func DemoChannel1() {
	pipe := make(chan string)

	go func() {
		for receiver := range pipe {
			fmt.Println(receiver)
		}
	}()

	pipe <- "water 1"
	pipe <- "water 2"
	pipe <- "water 3"
	close(pipe)
	time.Sleep(time.Millisecond)
}

func DemoChannel2() {
	pipe := make(chan string)
	done := make(chan bool)

	go func() {
		for {
			receiver, more := <-pipe
			fmt.Println(receiver)
			if !more {
				done <- false
				time.Sleep(time.Millisecond)
				fmt.Println("không in ra vì hàm đã thoát")
				return
			}
		}

	}()

	pipe <- "water 1"
	pipe <- "water 2"
	pipe <- "water 3"
	pipe <- "water 4"
	close(pipe)
	<-done
}
