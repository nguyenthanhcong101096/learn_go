package main

import "fmt"

// là 1 chuỗi các params
func addItem(item int, list ...int) {
	// 100 200 300 400 -> []int {100, 200, 300, 400}
	// []int {1,2,3,4} -> []int { []int {1,2,3,4}} -> error
	list = append(list, item)
	fmt.Println(list)
}

func changeItem(list ...int) {
	list[0] = 999
}

func main() {
	addItem(1, 100, 200, 300, 400)

	var list = []int {1, 2, 3, 4}
	addItem(100, list...)

	changeItem(list...)
	fmt.Println(list)
}