package main

func main() {
	//DemoSum()
	//DemoChannelBuffer5()
	//DemoChannel()
	/*sub := make(chan interface{})

	go func(c <-chan interface{}) {
		for data := range c {
			fmt.Println(data)
		}
	}(sub)

	sub <- "Hey there"*/

	//DemoChannelString()
	DemoChannel2()

}
