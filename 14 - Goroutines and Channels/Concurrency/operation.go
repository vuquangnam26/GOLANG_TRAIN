package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 3)
	for category, group := range data {
		go group.TotalPrice(category, channel)

	}
	time.Sleep(time.Second * 5)
	fmt.Println("--starting to reivec from channel", 1)
	for i := 0; i < len(data); i++ {
		fmt.Println("-- channel read pending", len(channel), "items in buffer, size", cap(channel))
		value := <-channel
		fmt.Println("--channel read complete 3", value)
		storeTotal += value
		time.Sleep(time.Second)
	}
	fmt.Println("Total:", ToCurency(storeTotal))
}
func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {

		total += p.Price
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "Channel sending 4", ToCurency(total))
	resultChannel <- total
	fmt.Println(category, "Channel send completed 5")
}
