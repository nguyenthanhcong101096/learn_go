package main

import (
	api "go_rest_api/api"
	models "go_rest_api/model"
	repo "go_rest_api/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	userN := models.User{Id: "1", Email: "nguyenthanhcong"}

	repo.CreateUser(&userN)

	router.HandleFunc("/api/v1/user/find", api.FindUser).Methods("GET")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err)
	}
}
