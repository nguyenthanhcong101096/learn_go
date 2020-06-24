package main

import "fmt"

//prints to stdout and puts an int on channel
func printHello(ch chan int) {
	fmt.Println("Hello from printHello")
	//send a value on channel
	ch <- 2
}

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Hello inline")

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go printHello(ch)

	for {
		select {
		case i := <-ch:
			fmt.Println("Recieved ", i)
		}
	}
}
