package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Products struct {
	name, category string
	price          float64
}

// Anonymus structsType
func AnonymusStructs() {

	prod := Products{name: "Kayak", category: "Watersports", price: 1500.00}
	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductsName  string
		ProductsPrice float64
	}{ProductsName: prod.name, ProductsPrice: prod.price})
	fmt.Println(builder.String())
	p2 := &prod
	p2.name = "OriginKayak"
	fmt.Println((*p2).name)
	fmt.Println(prod.name)
}
func clac(product *Products) {
	if product.price != 0 {
		product.price += product.price * 0.2
	}
}
func calcPointer(product *Products) *Products {
	if product.price != 0 {
		product.price += product.price * 0.2
	}
	return product
}
func CountructerPointer() {
	product := [2]*Products{
		NewProduct("kayak", "Watersports", 1500.00),
		NewProduct("LifeJacket", "Water", 355),
	}
	for _, pro := range product {
		fmt.Println(pro.name)
	}
}
func NewProduct(name, category string, price float64) *Products {
	return &Products{name: name, category: category, price: price}
}

type Supllier struct {
	name, city string
}
type product struct {
	name, category string
	price          float64
	*Supllier
}

func newProd(name, category string, price float64, supllier *Supllier) *product {
	return &product{name, category, price - 10, supllier}
}
func copyProdcut(product *product) product {
	p := *product
	s := *product.Supllier
	p.Supllier = &s
	return p
}
func main() {
	AnonymusStructs()
	kayak := Products{name: "Kayak", price: 1500.00}
	clac(&kayak)
	fmt.Println("name", kayak.name, "price", kayak.price)
	//direct
	kayak2 := &Products{name: "Kayak", price: 150.00}
	clac(kayak2)
	fmt.Println("name", kayak2.name, "price", kayak2.price)
	kayak3 := calcPointer(&Products{name: "Kayak", category: "Water", price: 15.00})
	fmt.Println("name", kayak3.name, "price", kayak3.price)
	CountructerPointer()
	//
	ace := &Supllier{"HN", "HN"}
	prod := [2]*product{newProd("kayak", "Water", 233, ace),
		newProd("LifeJacket", "Water", 222, ace)}
	for _, pro := range prod {
		fmt.Println(pro.name)
	}
}
