package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Shahriar-shudip/golang-microservies-tuitorial/product-api/data"
	"github.com/gorilla/mux"
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

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // ret
	id, _ := strconv.Atoi(vars["id"])
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	err := data.UpdateProduct(id, prod)
	if err == data.ErrProduct {
		http.Error(w, "unable to find the product ", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *Products) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJson(r.Body)

		if err != nil {
			http.Error(w, "unable to parse", http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		p.l.Println(ctx.Value(KeyProduct{}))
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})

}
