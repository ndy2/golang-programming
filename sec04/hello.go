package main

import "fmt"

type hotdog int

func main() {
	var a hotdog = 1
	var b int = 2
	fmt.Println(a, b)     // "1 2"
	fmt.Printf("%T\n", a) // "main.hotdog
	fmt.Printf("%T\n", b) // "int"
	// a = b // Cannot use 'b' (type int) as the type hotdog
	b = int(a)        // convert hotdog to int
	a = hotdog(b)     // convert int to hotdog
	fmt.Println(a, b) // "1 1"
}
