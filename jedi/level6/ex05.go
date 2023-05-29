package main

import "fmt"

type float float64
const pi float = 3.141592

type square struct {
	width  float
	height float
}

type circle struct {
	radius float
}

func (s square) getArea() float {
	return float(s.width * s.height)
}

func (c circle) getArea() float {
	return pi * c.radius * c.radius
}

type shape interface {
	getArea() float
}

func main() {
	fmt.Println("area of square : ", square{4, 5}.getArea()) // "area of square :  20"
	fmt.Println("area of circle: ", circle{4}.getArea()) // "area of circle:  50.265472"
}
