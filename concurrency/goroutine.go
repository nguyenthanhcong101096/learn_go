package main

import (
	"fmt"
	"time"
	"sync"
)

//Concurrency
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

// Goroutines
func g1() {
	fmt.Println("G1")
	wg.Done()
}

func g2() {
	fmt.Println("G2")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// var c chan string = make(chan string)

	// go pinger(c)
	// go printer(c)

	// var input string
	// fmt.Scanln(&input)

	fmt.Println("Start")

	wg.Add(2)

	go g1()
	go g2()

	wg.Wait()

	fmt.Println("End")
}