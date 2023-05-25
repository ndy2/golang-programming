package main

import "fmt"

type typeX int

var varY typeX

func main() {

	fmt.Println(varY)
	fmt.Printf("%T\n", varY)

	varY = 40
	fmt.Println(varY)

	convertedY := int(varY)
	fmt.Println(convertedY)
}
