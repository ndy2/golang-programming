package main

import "fmt"

type person struct {
	name string
}

func changeName(p *person, newName string) {
	(*p).name = newName
}

func main() {
	p := person{"haha"}
	changeName(&p, "papa")

	fmt.Println(p) // "{papa}"
}
