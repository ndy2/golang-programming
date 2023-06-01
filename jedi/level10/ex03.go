package main

import "fmt"

// starting with below code,
// pull the values off the channel using a for range loop

//func main() {
//	c := gen()
//	receive(c)
//
//	fmt.Println("about to exit")
//}
//
//func gen() <-chan int {
//	c := make(chan int)
//
//	for i := 0; i < 100; i++ {
//		c <- i
//	}
//
//	return c
//}

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
