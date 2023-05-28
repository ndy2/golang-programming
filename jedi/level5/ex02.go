package main

import "fmt"

type person struct {
	first                 string
	last                  string
	favoriteIcreamFlavors []string
}

func main() {
	haha := person{"nam", "deukyun", []string{"strawberry", "vanila"}}
	papa := person{"kim", "haha", []string{"choco", "vanila"}}

	m := map[string]person{
		"nam": haha,
		"kim": papa,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
