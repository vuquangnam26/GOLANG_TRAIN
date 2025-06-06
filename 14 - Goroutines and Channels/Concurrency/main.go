package main

import (
	"fmt"
	"time"
)

func enumerateProduct(chanel1, chanel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case chanel1 <- p:
			fmt.Println("Sent Product via chanel1", p.Name)
		case chanel2 <- p:
			fmt.Println("Sent Product via chanel 2", p.Name)

		}
	}
	close(chanel1)
	close(chanel2)
}
func InspectBuffer() {
	disPatchChannel := make(chan DispatchNotification, 100)
	//var sendOnlyChannel chan<- DispatchNotification = disPatchChannel
	//var reciveOnlyChannel <-chan DispatchNotification = disPatchChannel
	//or

	go DispatchOrder(chan<- DispatchNotification(disPatchChannel))
	//reciveDispatch((<-chan DispatchNotification)(disPatchChannel))
	//for {
	//	//	details := <-disPatchChannel
	//	//	fmt.Println("Dispatch to ", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	//	//}
	//	if details, open := <-disPatchChannel; open {
	//		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
	//			"x", details.Product.Name)
	//	} else {
	//		fmt.Println("Channel has been closed")
	//		break
	//	}
	//}
	productChanel1 := make(chan *Product)
	productChanel2 := make(chan *Product)
	go enumerateProduct(productChanel1, productChanel2)
	openChanels := 2
	for {
		select {
		case details, ok := <-disPatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("Dispatch channel closed")
				disPatchChannel = nil
				openChanels--
			}
		case product, ok := <-productChanel1:
			if ok {
				fmt.Println("Product", product.Name)
			} else {
				fmt.Println("Product channel closed")
				productChanel1 = nil
				openChanels--
			}
		default:
			if openChanels == 0 {
				goto alldone

			}
			fmt.Println("--No message ready to received --")
			time.Sleep(time.Millisecond * 500)
		}

	}
alldone:
	fmt.Println("All Value recevied")
}
func Clac() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	fmt.Println("main function complete")
}
func reciveDispatch(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}
func sendWithoutBlocking() {
	c1 := make(chan *Product, 2)
	c2 := make(chan *Product, 2)
	go enumerateProduct(c1, c2)
	time.Sleep(time.Second)
	for p := range c1 {
		fmt.Println("channel1 Recive Product", p.Name)
	}
	for p := range c2 {
		fmt.Println("channel 2 Recive Product", p.Name)
	}
}
func main() {
	//Clac()
	//InspectBuffer()
	sendWithoutBlocking()
}
