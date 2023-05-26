package main

import "fmt"

const (
	cur   = 2023
	prev1 = cur - iota
	prev2 = cur - iota
	prev3 = cur - iota
	prev4 = cur - iota
)

func main() {
	fmt.Println(
		cur, prev1, prev2, prev3, prev4,
	)
}
