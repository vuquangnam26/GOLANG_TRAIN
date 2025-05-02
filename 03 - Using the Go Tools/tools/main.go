// pagekage main
package main

import (
	"fmt"
)

func main() {
	printHello()
	for i := 0; i < 10; i++ {
		printNumber(i)
	}
}
func printHello() {
	fmt.Println("Hello World")
}
func printNumber(number int) {
	fmt.Println(number)
}
