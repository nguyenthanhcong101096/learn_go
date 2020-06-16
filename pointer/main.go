package main

import "fmt"

func main() {
	var pointer *int
	a := 100
	pointer = &a
	fmt.Println(pointer)      // in dịa chỉ

	// pointer thay đổi value a
	*pointer = 999            // <=> a:= 999
	fmt.Println(a)

	b := 123
	p2 := new(int)           // <=> var p2 *int
	p2 = &b
	fmt.Println(p2)

	//array pointer
	myArray := [3]int {1,2,3}
	var myPointer *[3]int
	myPointer = &myArray
	fmt.Println(myPointer)
	
	c := 888
	var myPointer1 *int = &c
	applyPointer(myPointer1)
	fmt.Println(c)
}

// truyền con trỏ vào hàm
func applyPointer(pointer *int) {
	*pointer = 777
}