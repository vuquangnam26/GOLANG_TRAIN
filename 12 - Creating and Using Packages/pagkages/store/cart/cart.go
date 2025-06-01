package cart

import "packages/store"

type Cart struct {
	CustomerName string
	Products     []store.Product
}

func (cart *Cart) GetTotal() (total float64) {
	for _, product := range cart.Products {
		total += product.Price()
	}
	return
}
