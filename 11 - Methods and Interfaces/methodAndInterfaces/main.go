package main

import "fmt"

// define interface
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func clacTotal(expese []Expense) (total float64) {
	for _, exp := range expese {
		total += exp.getCost(true)
	}
	return
}

type Person struct {
	name, city string
}

func processItem(item interface{}) {
	switch value := item.(type) {
	case Product:
		fmt.Println("Product:", value.name, "Price:", value.price)
	case *Product:
		fmt.Println("Product Pointer:", value.name, "Price:", value.price)
	case Service:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
	case Person:
		fmt.Println("Person:", value.name, "City:", value.city)
	case *Person:
		fmt.Println("Person Pointer:", value.name, "City:", value.city)
	case string, bool, int:
		fmt.Println("Built-in type:", value)
	default:
		fmt.Println("Default:", value)
	}
}
func processItem2(items ...interface{}) {
	for _, item := range items {

		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
}
func main() {
	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50, []string{}}
	fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	fmt.Println("Service:", insurance.description, "Price:", insurance.monthlyFee*float64(insurance.durationMonths))
	// using interface
	expense := []Expense{
		&Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50, []string{}},
	}
	var ex Expense = &kayak
	kayak.price = 100

	fmt.Println("Product field value:", kayak.price)
	fmt.Println("Expense method result:", ex.getCost(false))
	fmt.Println(clacTotal(expense))

	////Comparing
	//var e1 Expense = &Product{name: "Kayak"}
	//var e2 Expense = &Product{name: "Kayak"}
	//
	//var e3 Expense = Service{description: "Boat Cover"}
	//var e4 Expense = Service{description: "Boat Cover"}
	//fmt.Println("e1 == e2", e1 == e2)
	//fmt.Println("e3 == e4", e3 == e4)
	////Type Assertion
	typeAssertion := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Boat Cover", 12, 23, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

	for _, expense := range typeAssertion {
		s, ok := expense.(Service)
		if ok {
			fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))

		} else {
			fmt.Println("Expese", expense.getName())
		}
	}
	//Switch Type
	for _, expense := range typeAssertion {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case *Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		default:
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}

	//Empty InterFace
	data := []interface{}{
		expense, Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true}
	for _, item := range data {
		processItem(item)
	}
	processItem2(data...)

}
