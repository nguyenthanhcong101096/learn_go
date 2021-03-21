package main

import "fmt"

func main() {
	// Array là value type chứ không phải là ref tyoe
	// khai bao array
	var myArray [4]int
	fmt.Println(myArray)

	// gán giá trị cho array
	myArray[0] = 1000
	fmt.Println(myArray)

	// khai báo 1 array có khởi tạo giá trị
	arr := [3]int {1, 2, 3}
	fmt.Println(len(arr))

	// khai báo mảng không cần size
	phones := [...]string {"IPHONE", "SAMSUNG"}
	fmt.Println(phones[1])

	// array là value type không phải là ref type
	countries := [...]string {"VN", "JP", "EN", "UK"}
	copyCountries := countries

	copyCountries[0] = "CANADA"

	fmt.Println(countries)
	fmt.Println(copyCountries)

	// loop array
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}

	// blank identifier
	for index, _ := range countries {
		fmt.Println(index)
	}

	// mang 2 chieu
	matrix := [4][2]int {
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}

	fmt.Println(matrix)
}
