package api

import (
	"net/http"
	"restful_api/data"
)

var (
	books = data.Books{}
)

func init() {
	books.Initialize()
}

func Index(rw http.ResponseWriter, r *http.Request) {
	stores := books.All()
	err := stores.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Can not convert to JSON", http.StatusInternalServerError)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
	}
}
