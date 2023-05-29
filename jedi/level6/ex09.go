package main

import "fmt"

func main() {
	var result = applier(func(a int) int { return a + 1 }, 4)
	fmt.Println(result) // "5"
}

func applier(op func(int) int, a int) int {
	return op(a)
}
