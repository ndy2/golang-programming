package main

import "fmt"

func main() {
	year := 1997
	for {
		fmt.Println(year)
		if year == 2023 {
			break
		}
		year++
	}
}
