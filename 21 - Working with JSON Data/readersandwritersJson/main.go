package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{Product: &Kayak, Discount: 10.50}
	namdedItems := []Named{&dp, &Person{PersonName: "John Doe"}}
	encoder.Encode(namdedItems)
	fmt.Print(writer.String())
}
