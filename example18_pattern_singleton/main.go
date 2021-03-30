package main

import (
	"singleton/singleton"
	"fmt"
)

func main(){
	instance := singleton.GetInstance()
	fmt.Println(instance.AddItem())
	fmt.Println(instance.AddItem())
}