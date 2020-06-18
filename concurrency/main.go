// cách khai báo và cách sử dụng goroutine

package main

import (
	"fmt"
	"sync"
	"time"
)

func g1() {
	for {
		fmt.Println("AAA")
		time.Sleep(time.Second * 1)
	}
	wg.Done() // tín hiệu cho bik goroutines xong
}

func g2() {
	for {
		fmt.Println("BBB")
		time.Sleep(time.Second * 2)
	}
	wg.Done() // tín hiệu cho bik goroutines xong
}

var wg sync.WaitGroup // nơi chứa nhóm gorouties ta muốn controll synchronized

func main() {
	// go function_name()
	// Synchronized goroutines

	fmt.Println("Bắt đầu") // log1

	wg.Add(2) // tạo group 2 gorouties

	go g1() // g1()
	go g2() // g2()

	wg.Wait()

	fmt.Println("Kết thúc") // log2
}
