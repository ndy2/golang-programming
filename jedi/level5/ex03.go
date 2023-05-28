package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{vehicle{2, "white"}, true}
	s := sedan{vehicle{4, "black"}, true}

	fmt.Println(t) // "{{2 white} true}"
	fmt.Println(s) // "{{4 black} true}"
}
