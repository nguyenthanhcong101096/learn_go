package main

import (
	"fmt"
	stringHelper "hello-world/helper"
)

// trong projet chỉ nên có 1 file main.go

// pwd
// export GOPATH=/Users/congnt/Desktop/golang/workspace
// echo $GOPATH

func main() {
	fmt.Println("this is a main function")
	stringHelper.ConvertStringToInt()
}
