package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Decription string  `json:"description"`
	Price      float32 `json:"price"`
	SKU        string  `json:"sku"`
	CreateOn   string  `json:"-"`
	UpdateOn   string  `json:"-"`
	DeletedOn  string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

var productlist = []*Product{
	&Product{
		ID:         1,
		Name:       "shoe",
		Decription: "newly created",
		Price:      1.25,
		SKU:        "adf",
		CreateOn:   time.Now().UTC().String(),
		UpdateOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:         2,
		Name:       "Dress",
		Decription: "newly created",
		Price:      2.25,
		SKU:        "adf",
		CreateOn:   time.Now().UTC().String(),
		UpdateOn:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productlist
}
