package main

import (
	"reflect"
	"strings"
	//"strings"
	//"fmt"
)

func createPointerType(t reflect.Type) reflect.Type {
	return reflect.PointerTo(t)
}
func followPointerType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

var stringPtrType = reflect.TypeOf((*string)(nil))

func transformString(val interface{}) {
	elemValue := reflect.ValueOf(val)
	if elemValue.Type() == stringPtrType {
		upperValue := strings.ToUpper(elemValue.String())
		if elemValue.CanSet() {
			elemValue.SetString(upperValue)
		} else {
			Printfln("Cannot set value: %v", elemValue)
		}
	}
}

// array
func checkElemType(val interface{}, arrOrSliece interface{}) bool {
	elemType := reflect.TypeOf(val)
	arrOrSlieceType := reflect.TypeOf(arrOrSliece)
	return (arrOrSlieceType.Kind() == reflect.Array || arrOrSlieceType.Kind() == reflect.Slice) &&
		arrOrSlieceType.Elem() == elemType
}
func setValue(arrayOrSlice interface{}, index int, replacement interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	replacementVal := reflect.ValueOf(replacement)
	if arrayOrSliceVal.Kind() == reflect.Slice {
		elemVal := arrayOrSliceVal.Index(index)
		if elemVal.CanSet() {
			elemVal.Set(replacementVal)
		}
	} else if arrayOrSliceVal.Kind() == reflect.Ptr && arrayOrSliceVal.Elem().Kind() == reflect.Array && arrayOrSliceVal.Elem().CanSet() {
		arrayOrSliceVal.Elem().Index(index).Set(replacementVal)
	}

}
func main() {
	name := "John Doe"
	city := "London"
	hobby := "Running"
	slice := []string{name, city, hobby}
	array := [3]string{name, city, hobby}
	// Printfln("Slice (string): %v", checkElemType("testString", slice))
	// Printfln("Array (string): %v", checkElemType("testString", array))
	// Printfln("Array (int): %v", checkElemType(10, array))
	// t := reflect.TypeOf(name)
	// Printfln("Original Type: %v", t)
	// pt := createPointerType(t)
	// Printfln("Pointer Type: %v", pt)
	// Printfln("Follow pointer type: %v", followPointerType(pt))
	// Printfln("Original slice: %v", slice) 
	newCity := "Paris" 
	setValue(slice, 1, newCity)
	 Printfln("Modified slice: %v", slice)
	  Printfln("Original slice: %v", array) 
	  newCity = "Rome" 
	  setValue(&array, 1, newCity) 
	  Printfln("Modified slice: %v", array)

}
