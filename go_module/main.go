package main

import (
	"fmt"
	"go_module/nguoimexe"

	"github.com/leekchan/accounting"
)

// go mod init go_module
// go build
// ./go_module


// go mod vendor
// go mod graph

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))
	nguoimexe.Println()
}
