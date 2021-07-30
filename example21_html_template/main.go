package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type Student struct {
	Age   int
	Email string
}

func main() {
	person := Student{
		Age:   25,
		Email: "hola",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", person)

	if err != nil {
		log.Fatal("NOT FOUND", err)
	}
}
