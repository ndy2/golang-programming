package main

import "fmt"

func main() {
	f := func(name string) {
		fmt.Println("Hello", name)
	}
	f("haha")
}
