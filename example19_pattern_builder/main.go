package main

import (
	"fmt"
	"builder/builder"
)

func main(){
	normalBuidler := builder.GetBuilder("normal")
	// iglooBuidler := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuidler)
	buildHouse := director.BuildHouse()

	fmt.Println(buildHouse.GetWindownType())
	fmt.Println(buildHouse.GetFloor())
}