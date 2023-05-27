package main

import "fmt"

func main() {
	for c := 'A'; c <= 'Z'; c++ {
		for i := 0; i < 3; i++ {
			fmt.Printf("%U '%c'\n", c, c)
		}
		fmt.Println()
	}
}
