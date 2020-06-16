package main

import (
	"fmt"
	"errors"
)

var (
	ErrBadRequest = errors.New("Bad Request")
	ErrPagesMove  = errors.New("Redirect to")
)

func webCall(b bool) error {
	if b {
		return ErrBadRequest
	}
	return ErrPagesMove
}

func main() {
	if err := webCall(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Occurred")
			return

		case ErrPagesMove:
			fmt.Println("The Page moved")
			return

		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Life is good")
}