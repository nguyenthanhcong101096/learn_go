package main

import "fmt"

// khai báo struct
type Student struct {
	id int
	name string
}

func main() {
	//named
	st1 := Student {id: 123, name: "Congttl",}
	fmt.Println(st1.name)

	st2 := Student{456, "hola"}
	fmt.Println(st2.name)

	var st3 Student = struct {
		id int
		name string
	}{777, "holo"}
	fmt.Println(st3.name)

	//struct anonymous
	var anonymous = struct {
		email string
		age int
	} {
		email: "congttl@gmail.com",
		age: 25,
	}

	fmt.Println(anonymous)

	//pointer -> struct
	pointer := &Student{
		id: 999,
		name: "no name",
	}

	fmt.Println(&pointer)
	fmt.Println((*pointer).id) // pointer get value
	fmt.Println(pointer.name)  // pointer get value

	//anonymous filed -> k cần đặt tên cho filed
	type NoName struct {
		string
		int
	}

	noname := NoName{"hola", 777}
	fmt.Println(noname)

	// Struct in struct -> nested struct
	type Info struct {
		email string
		address string
	}

	type Person struct {
		id int
		name string
		info Info
	}

	person := Person{
		id: 123,
		name: "nguỹen thanh cong",
		info: Info{"cong@gmail.com", "Khánh hoà"},
	}

	fmt.Println(person.info.email)

	// so sánh 2 struct
	type Struct1 struct {
    id int
	}

	s1 := Struct1{1}
	s2 := Struct1{1}

	if s1 == s2 {
		fmt.Println("==")
	}
}