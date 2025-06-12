package main

import "fmt"

func Printfln(tempalte string, values ...interface{}) {
	fmt.Printf(tempalte+"\n", values...)
}
