package main

import "fmt"

func main() {
	fmt.Println("line6")
	defer fmt.Println("line7")
	fmt.Println("line8")
	defer fmt.Println("line9")

	/*
		실행 결과
		line6
		line8
		line9
		line7
	*/
}
