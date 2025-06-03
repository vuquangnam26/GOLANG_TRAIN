package main

import (
	"Composition/store"
	"fmt"
)

func main() {
	kayak := store.NewProduct("Kayak", "WaterSport", 268)
	lifeJacket := &store.Product{Name: "Life Jacket", Category: "WaterSport"}
	for _, p := range []*store.Product{kayak, lifeJacket} {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}
	//Composing type
	boats := []*store.Boat{store.NewBoat("Kayak", 275, 1, false), store.NewBoat("Canoe", 400, 3, false), store.NewBoat("Tender", 650.25, 2, true)}
	for _, b := range boats {
		fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Price(0.2))
	}

	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	Name, price, Price := deal.GetDetail()
	fmt.Println("Name:", Name)
	fmt.Println("Price field:", price)
	fmt.Println("Price method:", Price)
	//type OfferBundle struct {
	//	*store.SpecialDeal
	//	*store.Product
	//}
	//bundle := OfferBundle{
	//	store.NewSpecialDeal("Weekend Special", kayak, 50),
	//	store.NewProduct("Lifrejacket", "Watersports", 48.95)}
	//fmt.Println("Price:", bundle.Price(0))

	products := map[string]store.ItemForSale{"Kayak": store.NewBoat("Kayak", 279, 1, false), "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50)}
	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("Name:", item.GetName(), "Category:", item.GetCategory(), "Price:", item.(store.ItemForSale).Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}
}
