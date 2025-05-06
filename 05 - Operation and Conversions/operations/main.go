package main

import (
	"fmt"
	"strconv"
)

func CompareingPointer() {
	fist := 100
	second := &fist
	third := &fist
	alpha := 100
	beta := &alpha
	fmt.Println(second == third)
	fmt.Println(second == beta)
	fmt.Println(*second == *beta)
	fmt.Println(*second == *third)
}
func LogicalOperation() {
	maxMph := 50
	passengerCapacity := 100
	airbags := true
	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar
	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)
}
func Operation() {
	fmt.Println("Hello World")
	price, tax := 275.00, 27.40
	sum := price + tax
	difference := price - tax
	product := sum * difference
	qoutient := price / tax
	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(qoutient)
}
func ExplicitType() {
	kayak := 275
	soccerBall := 19.50
	total := float64(kayak) + soccerBall
	fmt.Println(total)
}
func ExplicitParseInt() {
	val1 := "1000"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parse Int", smallInt)
	} else {
		fmt.Println("Could not parse Int", val1, int1err)
	}
}
func ParsePoitner() {
	val1 := "4.895e+01"
	float1, float1err := strconv.ParseFloat(val1, 64)
	if float1err == nil {
		fmt.Println("Parse Float", float1)
	} else {
		fmt.Println("Could not parse Float", val1, float1err)
	}
}
func main() {
	Operation()
	CompareingPointer()
	LogicalOperation()
	ExplicitType()
	ExplicitParseInt()
	ParsePoitner()
}
