package main

import (
	"fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())

	// usinmg nested packages
	cart := cart.Cart{CustomerName: "alice", Products: []store.Product{*product}}
	fmt.Println(cart)
	fmt.Println("Products Tottal:", cart.GetTotal())
}
