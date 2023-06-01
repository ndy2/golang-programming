package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error : %v", ce.info)
}

func main() {
	ce := customErr{"err info"}
	fmt.Println(ce)        // "here is the error : err info"
	fmt.Printf("%T\n", ce) // "main.customErr"

	foo(ce) //"foo ran -  here is the error : err info err info"
}

func foo(e error) {
	fmt.Println("foo ran - ", e, e.(customErr).info)
}
