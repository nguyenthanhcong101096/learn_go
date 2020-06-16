package main

import "fmt"

type error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func New(text string) error {
	return errorString{text}
}

// webCall performs a web operation.
func webCall() error {
	return New("Bab Request")
}

func main() {
	
	if err := webCall(); err != nil {
		fmt.Println(err)
	
	}

	fmt.Println("Life is good")
}
