package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

func DispatchOrder(chanel chan<- DispatchNotification) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	orderCount := rand.Intn(5) + 5
	fmt.Println("orderCount", orderCount)
	for i := 0; i < orderCount; i++ {
		chanel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10) + 1,
			Product:  ProductList[rand.Intn(len(ProductList))],
		}
		time.Sleep(time.Millisecond * 750)
	}
	close(chanel)
}
