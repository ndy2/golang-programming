package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("I am %v %v. My age is %v\n", p.first, p.last, p.age)
}

func main() {
	person{"nam", "deukyun", 27}.speak()
}
