package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("hola.txt", "one.txt")

	if err != nil {
		log.Fatal("ParseFiles Error", err)
	}

	tpl.ExecuteTemplate(os.Stdout, "hola.txt", "25tuoi")
}
