package main

import "fmt"

// starting with below code,
// pull the values off the channel using a select statement
//
//func main() {
//	q := make(chan int)
//	c := gen2(q)
//
//	receive2(c)
//
//	fmt.Println("about to exit")
//}
//
//func gen2(q <-chan int) <-chan int {
//	c := make(chan int)
//
//	for i := 0; i < 10; i++ {
//		c <- i
//	}
//
//	return c
//}

func main() {
	q := make(chan int)
	c := gen2(q)

	receive2(c, q)
	fmt.Println("about to exit")
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive2(c <-chan int, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
