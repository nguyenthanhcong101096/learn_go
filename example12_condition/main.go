package main

import "fmt"

// Condition
func main() {
	//if else
	number := 10
	if number == 10 {
		fmt.Println("number == 10")
	} else {
		fmt.Println("number != 10")
	}

	if num := 100; num == 100 {
		fmt.Println("num == 100")
	} else {
		fmt.Println("num != 100")
	}

	// switch - case
	num1 := 10
	switch num1 {
	case 0, 1, 2:
		fmt.Println("number == 0, 1, 2")
	case 5:
		fmt.Println("number == 5")
	case 10:
		if number == 10 {
			fmt.Println("break here")
			break
		}

		fmt.Println("number == 10")
		fallthrough //tiếp tục chạy case bên dưới
	case 15:
		fmt.Println("number == 15")
		fallthrough
	case 20:
		if number == 20 {
			goto handleNumber
		}

		fmt.Println("number == 20")
	handleNumber:
		fmt.Println("handle for case 20")
	default:
		fmt.Println("unknow")
	}

	//or
	switch {
	case num1 > 10:
		fmt.Println("number > 10")
	case num1 < 10:
		fmt.Println("number < 10")
	default:
		fmt.Println("number == 10")
	}

	//loops
	// for init; condition; post

	//break; continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("out of loop")
			break // out of loop
		}

		if i == 2 {
			fmt.Println("skip index 2")
			continue // bỏ index 2
		}

		fmt.Println(i)
	}

	//while
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	for {
		fmt.Println(" infinite loop")
	}
}
