package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u0 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u1 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u2 := user2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var users []user2
	err := json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println(err)
	}

	println(reflect.DeepEqual(users[0], u0)) // "true"
	println(reflect.DeepEqual(users[1], u1)) // "true"
	println(reflect.DeepEqual(users[2], u2)) // "true"
}
