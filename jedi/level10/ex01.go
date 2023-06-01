package main

import "fmt"

// make this code works!
//
//func main() {
//	c := make(chan int)
//	c <- 42
//	fmt.Println(<-c)
//}

func main() {
	// solution 1 - using func literal
	//c := make(chan int)
	//go func() { c <- 42 }()

	// solution 2 - using buffer
	c := make(chan int, 1)
	c <- 42

	fmt.Println(<-c)
}
