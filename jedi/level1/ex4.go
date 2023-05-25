package main

import "fmt"

type hotdog int

var hello hotdog

func main() {
	fmt.Println(hello)
	fmt.Printf("%T\n", hello)
}
