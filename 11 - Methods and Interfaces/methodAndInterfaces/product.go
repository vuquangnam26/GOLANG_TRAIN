package main

type Product struct {
	name, category string
	price          float64
}

// implement
func (p *Product) getName() string {
	return p.name
}
func (p *Product) getCost(_ bool) float64 {
	return p.price
}
