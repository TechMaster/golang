package main

import "fmt"

func DemoChannelBuffer1() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func DemoChannelBuffer2() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 //Hỏi: Tại sao lỗi vậy?
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func DemoChannelBuffer3() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3 //Lệnh này không bị lỗi
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
func DemoChannelBuffer4() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	for i := range ch {
		fmt.Println(i) //fatal error: all goroutines are asleep - deadlock!
	}
}

func DemoChannelBuffer5() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch) //Thêm dòng này chạy ok
	for i := range ch {
		fmt.Println(i)
	}
}
