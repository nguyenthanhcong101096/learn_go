package main

import "fmt"

type Error interface {
	messages()
}

type Standard struct {
	code    int
	message string
}

func (a Standard) messages() {
	fmt.Println(a.code)
	fmt.Println(a.message)
}

func new(code int, message string) Error {
	return Standard{code, message}
}

var (
	RecordNotFound  = new(404, "ActiveRecord::RecordNotFound")
	RecordNotUnique = new(409, "ActiveRecord::RecordNotUnique")
)

// create user

func createUser(b bool) Error {
	if b {
		return RecordNotFound
	}

	return RecordNotUnique
}

func main() {
	if err := createUser(true); err != nil {
		err.messages()
	}
}
