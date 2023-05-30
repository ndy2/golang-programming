package main

import (
	"fmt"
	"sort"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []person{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	for _, p := range users {
		sort.Strings(p.Sayings)
	}
	less := func(i, j int) bool {
		if users[i].Age == users[j].Age {
			return users[i].Last < users[j].Last
		} else {
			return users[i].Age < users[j].Age
		}
	}
	sort.Slice(users, less)

}
