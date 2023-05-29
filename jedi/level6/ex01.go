package main

import "fmt"

func main() {
	fmt.Println(foo()) // "100"
	fmt.Println(bar()) // "100 bar"
}

func foo() int {
	return 100
}

func bar() (int, string) {
	return 100, "bar"
}
