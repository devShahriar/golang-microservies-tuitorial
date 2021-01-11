package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	if r.Method == http.MethodPut {
		regex := regexp.MustCompile(`/([0-9]+)`)
		p.l.Printf("%v", r.URL.RawPath)
		g := regex.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Println(r.URL.Path)
		p.l.Println(g)
		if len(g) != 1 {
			http.Error(w, "invalid", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(w, "invalid", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)

		p.l.Println("got id ", id)
		p.updateProduct(id, w, r)
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

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {

	prod := &data.Product{}
	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(w, "unable to parse", http.StatusInternalServerError)
	}
	p.l.Printf("Prod %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
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
