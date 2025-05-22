package main

import (
	"fmt"
	"reflect"
	"sort"
)

func CopySlice() {
	products := []string{"Kayak", "LifeJacket", "Paddle", "Hat"}
	allNames := products[1:]
	someNames := make([]string, 2)
	copy(someNames, allNames)
	//
	//// neu khong initlize slice thi copy se khong thuc thi duoc do cap khong duoc dinh nghia
	////specifying ranges when copying slices
	someNames = []string{"Boots", "Canoe"}
	copy(someNames[1:], allNames[2:3])
	fmt.Println("SomeName", someNames)
	//fmt.Println("All", allNames)
	//// copying Slices with Different sizes
	replacementSlices := []string{"Canoe", "Boats"}
	copy(products, replacementSlices)
	//fmt.Println("products", products)
	// delete
	deleted := append(products[:2], products[3:]...)
	fmt.Println("delte", deleted)
	fmt.Println("product", products)
	sort.Strings(products)
	for index, value := range products[2:] {

		fmt.Println("value", value, "index", index)

	}
	p2 := products
	fmt.Println("Equal", reflect.DeepEqual(products, p2))
	arrPtr := (*[3]string)(products)
	array := arrPtr
	fmt.Println(array)
}
func Map() {
	products := map[string]float64{"Kayyak": 279, "LifeJacket": 95.54, "Hat": 0}
	delete(products, "Hat")

	if value, ok := products["Hat"]; ok {
		fmt.Println("Hat", value)
	} else {
		fmt.Println("Hat is empty")
	}
	for key, value := range products {
		fmt.Println("Key:", key, "Value:", value)
	}
	//enumeratinga map in order
	keys := make([]string, 0, len(products))
	for key, _ := range products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", products[key])
	}

}
func String() {
	var price = "€48.95"
	for index, char := range price {
		fmt.Println(index, char, string(char))
	}
	//byte
	for index, char := range []byte(price) {
		fmt.Println(index, char)
	}
}
func Slice() {
	names := []string{"Kayak", "LifeJacket", "Car"}
	names = append(names, "Boat", "Cycle")
	fmt.Println(names)
	moreNames := []string{"Halloween Chrismarst"}
	appendNames := append(names, moreNames...)
	fmt.Println(appendNames)

	//create Slice form exits array
	someSlice := appendNames[1:4]
	fmt.Println(someSlice)
	allSlice := appendNames[:]
	fmt.Println(allSlice)

	//making slice predictable
	products := [4]string{"Kayak", "LifeJacket", "Paddle", "Hat"}
	someProducts := products[1:3:3]
	allProducts := products[:]
	someProducts = append(someProducts, "Gloves")
	someProducts = append(someProducts, "Boots")
	fmt.Println(someProducts)
	fmt.Println("len", len(someProducts), "cap", cap(someProducts))
	fmt.Println(allProducts)
	fmt.Println("len", len(allProducts), "cap", cap(allProducts))
	//creating slice from other slice
	allProd := products[1:]
	someProd := allProd[1:3]
	allProd = append(allProd, "Gloves")
	allProd[1] = "Canoe"
	fmt.Println("SomeProd", someProd, "cap", cap(someProd))
	fmt.Println("AllProd", allProd, "cap", cap(allProd))
}
func main() {
	//var names [3]string
	//names[0] = "Kayak"
	//names[1] = "LifeJacket"
	//names[2] = "Paddle"
	names := [3]string{"Kayak", "LifeJacket", "Paddle"}
	fmt.Println(names)
	// mutilDimensional
	var coords [3][3]int
	coords[1][2] = 10
	// assign
	otherArr := names
	names[0] = "Yubby"
	fmt.Println("names", names)
	fmt.Println("other", otherArr)
	// pointer
	otherArrPoniter := &names
	names[1] = "Diff"
	fmt.Println("names", names)
	fmt.Println("other", *otherArrPoniter)

	for index, value := range names {
		fmt.Println(index, value)
	}
	//Slice()
	//CopySlice()
	Map()
}
