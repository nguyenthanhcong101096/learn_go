package main

import (
	"html/template"
	"log"
	"net/http"
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
	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request) {
	person := Student{Age: 25, Email: "hola"}
	err := tpl.ExecuteTemplate(w, "index.gohtml", person)

	if err != nil {
		log.Fatal("NOT FOUND", err)
	}
}
