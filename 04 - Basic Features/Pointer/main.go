package main

import (
	"fmt"
	"sort"
)

func PointerSlice() {
	names := [3]string{"Alice", "Charlie", "Bob"}
	secondName := names[1]
	fmt.Println(secondName)
	fmt.Println(secondName)
	secondPosition := &names[1]
	fmt.Println(*secondPosition)
	sort.Strings(names[:])
	fmt.Println(*secondPosition)

}
func main() {
	fist := 100
	var second = &fist
	fist++
	*second++
	var myNewPointer *int
	myNewPointer = second
	*myNewPointer++
	fmt.Println("fist =", fist)
	fmt.Println("second = ", *second)
	PointerSlice()
}
