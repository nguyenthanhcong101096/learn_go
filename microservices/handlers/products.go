package handlers

import (
	"log"
	data "microservices/data"
	"net/http"
	"regexp"
	"strconv"
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

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)

	if len(g) != 1 || len(g[0]) != 2 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	idString := g[0][1]
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Khong the convert JSON", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ProductNotFound || err != nil {
		http.Error(rw, "Khong tim thay", http.StatusNotFound)
		return
	}
}
