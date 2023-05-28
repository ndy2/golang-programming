package main

import "fmt"

type person struct {
	first                 string
	last                  string
	favoriteIcreamFlavors []string
}

func main() {
	haha := person{"nam", "deukyun", []string{"strawberry", "vanila"}}
	fmt.Println(haha) // "{nam deukyun [strawberry vanila]}"
	for _, v := range haha.favoriteIcreamFlavors {
		fmt.Println(v)
	}
}
