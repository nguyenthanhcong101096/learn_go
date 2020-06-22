package main

import (
	"net/http"
	"restful_api/api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/books", api.Index).Methods(http.MethodGet)

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err)
	}
}
