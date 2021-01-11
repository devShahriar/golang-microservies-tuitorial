package handlers

import (
	"log"
	"net/http"

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

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJson(w)
	if err != nil {
		http.Error(w, "unable to parse", http.StatusInternalServerError)
	}
}
