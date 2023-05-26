package main

import "fmt"

func main() {
	lhs := 78
	rhs := 99
	g := lhs == rhs
	h := lhs <= rhs
	i := lhs >= rhs
	j := lhs != rhs
	k := lhs < rhs
	l := lhs > rhs

	fmt.Println(
		g, h, i, j, k, l,
	)
}
