package main

import (
	"fmt"
	"math/rand"
)

func go1(channel chan int) {
	for {
		channel <- rand.Intn(100)
		fmt.Println(rand.Intn(100))
	}
}

func go2(channel chan int) {
	for {
		value := <-channel
		if value == 2 {
			fmt.Println("Random 2")
			return
		}
	}
}

func main() {
	channel := make(chan int)
	go go1(channel)
	go go2(channel)

	var input string
	fmt.Scanln(&input)
}
