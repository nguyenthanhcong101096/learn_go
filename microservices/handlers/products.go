package handlers

import (
	"log"
	data "microservices/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET")

	lg := data.GetProducts()
	err := lg.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Can not convert to JSON", http.StatusInternalServerError)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Khong the convert JSON", http.StatusBadRequest)
	} else {
		data.AddProduct(prod)
		rw.WriteHeader(200)
		p.l.Println(prod)
	}
}
