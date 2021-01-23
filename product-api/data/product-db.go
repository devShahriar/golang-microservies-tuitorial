package data

import (
	"encoding/json"
	"fmt"
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

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
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

func AddProduct(p *Product) {
	p.ID = getNextID()
	productlist = append(productlist, p)
}

func getNextID() int {
	lp := productlist[len(productlist)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, prod *Product) error {
	_, pos, err := getProductById(id)
	if err != nil {
		return err
	}
	prod.ID = id
	productlist[pos] = prod
	return nil
}

var ErrProduct = fmt.Errorf("Product not found")

func getProductById(id int) (*Product, int, error) {
	for i, p := range productlist {
		if id == p.ID {
			return p, i, nil
		}

	}
	return nil, -1, ErrProduct
}
