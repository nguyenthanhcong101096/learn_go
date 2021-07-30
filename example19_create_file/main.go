package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// CREATE STRING
	// assign to variable
	name := "James"
	s := fmt.Sprint(`mas ` + name + `menos`)

	// CREATE FILE
	// io.Copy to the file
	nf, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}

	io.Copy(nf, strings.NewReader(s))
}
