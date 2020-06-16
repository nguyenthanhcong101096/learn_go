package main

import "fmt"

func main(){
	// khai báo slice
	var slice []int
	fmt.Println(slice)

	// khai báo và khởi tạo cho slice
	var slice1 = []int {1,2,3,4}
	fmt.Println(slice1)

	// tạo slice từ 1 array
	var array = [4]int {1,2,3,4}
	slice2 := array[1:3] // array[1] -> array[3-1=2] <=> array[2]
	fmt.Println(slice2)

	slice3 := array[:]
	fmt.Println(slice3)

	slice4 := array[2:] // index 2 -> đến cuối
	fmt.Println(slice4)

	slice5 := array[:3] // 0 -> index 3 - 1
	fmt.Println(slice5)

	// tạo slice từ 1 slice khác
	var slice6 = []int {1,2,3,4,5,6,7,8,9}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[1:] 
	fmt.Println(slice8)

	// slice => reference type
	var arr1  = [5]int {1,2,3,4,5}
	slice9 := arr1[:]
	slice9[0] = 9999
	fmt.Println(slice9)

	// length và capacity của slice
	// len: số lượng phần tử của slice
	// cap: số lượng phần tử của underlying array bắt đầu từ vị trí start khi mà slice được tạo
	countries := [...]string {"VN","JP","UK","USA"}
	slice10 := countries[2:3]
	fmt.Println(slice10)
	fmt.Println(len(slice10))
	fmt.Println(cap(slice10))

	// make, copy, append
	// make -> khai báo len 2 và cap 5, không khai báo cap thì cap = len
	slice11 := make([]int, 2, 5)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	// append
	var slice12 []int
	slice12 = append(slice12, 100)
	fmt.Println(slice12)

	//copy
	src := []string {"A", "B", "C", "D"}
	dest := make([]string, 2)

	copy(dest, src)
	fmt.Println(dest)
}