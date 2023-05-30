package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 10, 8, 3, -3, 5, -1}
	sort.Ints(a)

	fmt.Println(a) // "[-3 -1 3 4 5 8 10]"
}
