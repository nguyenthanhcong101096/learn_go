package main

import "fmt"

// Interface -> vehicle { car, motocyle }
// multiple interface

type Animal interface {
	speak()
}

type Move interface {
	move()
}

// embed interface
type NextAnimal interface {
	Move
	Animal
}

// empty interface
func goout(i interface{}) {
	fmt.Println(i)
}

type Dog struct {}

func (d Dog) speak() {
	fmt.Println("Gâu gâu")
}

func (d Dog) move() {
	fmt.Println("Chạy bằng 4 chân")
}

func main(){
	// khai báo interface
	var animal Animal
	animal = Dog {}
	animal.speak()

	// multiple interface
	dog := Dog{}

	var m Move = dog
	m.move()

	var a Animal = dog
	a.speak()

	// Embed interface
	dog1 := Dog{}

	var embed NextAnimal = dog1
	embed.move()
	embed.speak()

	// empty interface
	goout(10)
}