package main

import "fmt"

func main() {
	var x = []int{1, 2, 3, 4}
	fmt.Println(foo(x...)) // "10"
}

func foo(ints ...int) int {
	var result = 0
	for _, x := range ints {
		result += x
	}
	return result
}
