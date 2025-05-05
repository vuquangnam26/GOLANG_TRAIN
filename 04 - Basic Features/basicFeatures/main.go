package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Int())
	var price float32 = 275.00
	var tax float32 = 27.25
	const quantity = 100
	fmt.Println(quantity * (price + tax))

}
