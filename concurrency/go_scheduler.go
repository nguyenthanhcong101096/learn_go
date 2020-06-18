package main

import (
	"fmt"
	"runtime"
)

func g1() { fmt.Println("G1") }
func g2() { fmt.Println("G2") }
func g3() { fmt.Println("G3") }
func g4() { fmt.Println("G4") }

func main() {
	runtime.GOMAXPROCS(2) // khởi tạo 2 Process


}
