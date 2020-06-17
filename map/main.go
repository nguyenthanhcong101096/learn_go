package main

import "fmt"

func main() {
	var myMap = make(map[string]int)
	fmt.Println(myMap)

	var myMap1 map[string]int
	fmt.Println(myMap1)

	// khai báo với giá trị khởi tạo
	myMap2 := map[string]int{"cong": 1, "lien": 2}
	fmt.Println(myMap2)

	// add item
	myMap2["key1"] = 1
	myMap2["key2"] = 2
	fmt.Println(myMap2)

	// delete item
	delete(myMap2, "key1")
	fmt.Println(myMap2)

	// size map
	fmt.Println(len(myMap2))

	// Map là reference key
	myMap3 := myMap2
	delete(myMap3, "key2")
	fmt.Println(myMap2)

	// truy cập 1 phần tử trong map
	value, found := myMap2["family"]

	if found {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}

	// trong map không có toán tử ==
}
