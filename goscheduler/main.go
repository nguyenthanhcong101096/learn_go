package main

import (
	"fmt"
	"runtime"
)

func g1() { fmt.Println("G1") }
func g2() { 
	fmt.Println("G2") 

	// thuc thi logic 1
	// thuc thi logic 2
}

func g3() { fmt.Println("G3") }
func g4() { fmt.Println("G4") }

func main() {
	runtime.GOMAXPROCS(3)

	numberP := runtime.NumCPU()
	fmt.Println(numberP)

	numberP1 := runtime.GOMAXPROCS(0)
	fmt.Println(numberP1)
}