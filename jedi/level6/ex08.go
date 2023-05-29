package main

import "fmt"

func main() {
	fn := hello()
	fmt.Println(fn(5))      // "6"
	fmt.Println(hello()(6)) // "7"
}

func hello() func(int) string {

	return func(a int) string {
		return fmt.Sprintf("%v", a+1)
	}
}
