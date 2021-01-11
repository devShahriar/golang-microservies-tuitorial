package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mico/data"
)

//Products exported
type Products struct {
	l *log.Logger
}

//NewProducts exported
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJson(w)
	if err != nil {
		http.Error(w, "unable to parse", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {

	prod := &data.Product{}
	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(w, "unable to parse", http.StatusInternalServerError)
	}
	p.l.Printf("Prod %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	prod := &data.Product{}
	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(w, "unable to parse", http.StatusInternalServerError)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProduct {
		http.Error(w, "unable to find the product ", http.StatusBadRequest)
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}
