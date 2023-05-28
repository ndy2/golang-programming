package main

import "fmt"

func main() {
	s := struct {
		doors int
		color string
	}{4, "black"}

	fmt.Println(s) // "{4 black}"
}
