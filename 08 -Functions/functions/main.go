package main

import "fmt"

func printSuppliers(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Supplier", supplier)
	}
}
func calcTax(price float64) (float64, bool) {
	if price > 10 {
		return price * 10, true
	} else {
		return 0, false
	}
}
func calcTotalPrice(product map[string]float64, minSpeed float64) (total, tax float64) {
	fmt.Println("Calculating total price")
	defer fmt.Println("first defer")
	for _, price := range product {
		if taxAmount, due := calcTax(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	defer fmt.Println("second defer")
	fmt.Println("Calculating total price2")

	return
}

func main() {
	names := []string{"Accme Kayas", "Boots"}
	printSuppliers("Kayak", names...)
	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	products := map[string]float64{"Kayak": 275, "Lifeajacket": 48.95}
	for product, price := range products {
		var calcFunc func(float64) (float64, bool)
		if taxAmount, taxDue := calcTax(price); taxDue {
			calcFunc = calcTax
			total, _ := calcFunc(price)
			fmt.Println("Tax", taxAmount, "prodcut", product, "calc", total)
		} else {
			fmt.Println("no tax", "prodcut", product)
		}
	}
	total1, tax := calcTotalPrice(products, 10)
	fmt.Println("Total price", total1, "tax", tax)
}
