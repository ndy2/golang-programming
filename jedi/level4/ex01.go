package main

import "fmt"

func main() {
	arr := [6]int{9, 7, 0, 8, 2, 1}
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", arr)
}
