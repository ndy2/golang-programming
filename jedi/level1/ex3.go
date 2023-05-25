package main

import "fmt"

func main() {
	var a int = 42
	var b string = "James Bond"
	var c bool = true
	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s) // "42      James Bond      true"
}
