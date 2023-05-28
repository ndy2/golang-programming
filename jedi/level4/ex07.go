package main

import "fmt"

func main() {
	names := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James"},
	}

	for _, row := range names {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
}
