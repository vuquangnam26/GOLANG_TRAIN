package main

import (
	"fmt"
	"strconv"
)

func Initstate_FlowControl() {
	priceString := "275"
	if kayakPrice, err := strconv.Atoi(priceString); err != nil {
		fmt.Println("Price", kayakPrice)
	} else {
		fmt.Println("Error", err)
	}
}
func Loop() {
	for counter := 0; counter < 3; counter++ {
		if counter == 1 {
			continue
		}
		fmt.Println(counter)
	}
	product := "Kayak"
	for index, character := range product {
		// mac dinh la co index k muon dung thi thay the bang _
		fmt.Println("Index", index, "Character", character)
	}
}
func lableState() {
	counter := 0
target:
	fmt.Print("counter", counter)
	counter++
	if counter < 5 {
		goto target
	}
}
func Switch() {
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lower case is positions", index)
				break
			}
			fmt.Println("Uppercase is positions", index)
		case 'y':
			fmt.Println("Yes is positions", index)
		default:
			fmt.Println("Character", string(character), "at positon", index)
		}
	}
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime value", counter/2)
		default:
			fmt.Println("Not prime value", counter/2)
		}
	}
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value", val, counter)
		default:
			fmt.Println("Not prime value", val, counter)
		}

	}
}
func main() {
	//	Initstate_FlowControl()
	//	Loop()
	//Switch()
	lableState()
}
