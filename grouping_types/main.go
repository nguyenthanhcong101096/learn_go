package main

import "fmt"

// --------------------
// Grouping By Behavior
// --------------------

type Behavior interface {
	speak()
}

// -----------------
// Grouping By State
// -----------------
type Animal struct {
	Name string
}

type Cat struct {
	Animal     // 1 gôm nhóm 1 phần chung của struct, có thể sử dụng attribute của 1 phần chung
	PackFactor int
}

type Tiger struct {
	Animal      // 1 gôm nhóm 1 phần chung của struct, có thể sử dụng attribute của 1 phần chung
	ClimbFactor int
}

func (a Animal) speak() {
	fmt.Println("UGH!", "Tôi là", a.Name)
}

func (d Tiger) speak() {
	fmt.Println("Ghừ!", "Tôi là", d.Name, ", Tôi di chuyển bằng", d.ClimbFactor, "Và tôi sống ở rừng, tôi kêu Ghừ!")
}

func (c Cat) speak() {
	fmt.Println("Meow!", "Tôi là", c.Name, ", Tôi di chuyển bằng", c.PackFactor, "Và tôi kêu Meow!")
}

func main() {
	tiger := Tiger{Animal{"Tiger"}, 4}
	cat := Cat{Animal{"Cat"}, 4}

	behaviors := []Behavior{tiger, cat}

	for _, animal := range behaviors {
		animal.speak()
	}
}
