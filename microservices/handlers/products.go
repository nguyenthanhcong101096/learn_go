package handlers

import (
	"context"
	"log"
	data "microservices/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lg := data.GetProducts()
	err := lg.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Can not convert to JSON", http.StatusInternalServerError)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
	}
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
	rw.WriteHeader(200)
	p.l.Println(prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable convert to id", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, &prod)

	if err == data.ProductNotFound || err != nil {
		http.Error(rw, "Khong tim thay", http.StatusNotFound)
		return
	}
}

func (p Products) Profile(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"message": "My Profile"}`))
}

// middleware

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "Khong the convert JSON", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

var tokens = map[string]string{"user1": "user1", "user2": "user2", "user3": "user3"}

func (p Products) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if _, ok := tokens[token]; ok {
			next.ServeHTTP(rw, r)
		} else {
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusForbidden)
			rw.Write([]byte(`{"message": "Authenticate fail"}`))
		}
	})
}
