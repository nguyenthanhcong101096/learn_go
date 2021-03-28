package main

import "fmt"

// Variadic function bản chất nó là 1 func trong golang
// Nhận vào vô số tham số mà không cần giới hạn than số là bao nhiêu

func addItem(item int, list ...int) {
	// 100 200 300 -> slice int[] {100 200 300 }
	list = append(list, item)
}

func main() {
	addItem(1, 100,200,3000)

	list := []int {1,1,2,2}
	addItem(100, list...)
}