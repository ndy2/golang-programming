package main

import "fmt"

// write a program that
//
//	launch 10 goroutines - each goroutine adds 10 numbers to a channel
//	pull the numbers off the channel and print them
func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			for j := 0; j < 10; j++ {
				c <- i*10 + j
			}
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}
