package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak(word string) {
	fmt.Println("person", p.name, "speaks : ", word)
}

type human interface {
	speak(word string)
}

func saySomething(h human, word string) {
	h.speak(word)a
}

func main() {
	var p = person{"haha", 27}
	saySomething(&p, "hello golang")
	// saySomething(p, "hello golang") // not working!!!
	//  cannot use p (variable of type person) as human value in argument to saySomething: 
	// person does not implement human (method speak has pointer receiver)
}
