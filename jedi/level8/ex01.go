package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{"James", 32}
	u2 := user{"Moneypenny", 27}
	u3 := user{"M", 54}

	users := []user{u1, u2, u3}
	fmt.Println(users) // "[{James 32} {Moneypenny 27} {M 54}]"

	bs, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) // "[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]"
}
