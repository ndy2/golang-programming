package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("dec : %d\n", a)
	fmt.Printf("bin : %b\n", a)
	fmt.Printf("hex : %x\n", a)

	a <<= 1
	fmt.Printf("dec : %d\n", a)
	fmt.Printf("bin : %b\n", a)
	fmt.Printf("hex : %x\n", a)
}
