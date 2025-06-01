package store

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}
func (p *Product) Price() float64 {
	return p.price
}
func (p *Product) SetPrice(newPrices float64) {
	p.price = newPrices
}
