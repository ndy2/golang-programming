package main

import "fmt"

func main() {
	favoriteSongMap := map[string][]string{
		`haha`: []string{`song1`, "song2"},
		`papa`: []string{`song3`, "song4"},
	}

	favoriteSongMap["new jeans"] = []string{"hype boy", "attention"}

	fmt.Println(favoriteSongMap)

	for k, v := range favoriteSongMap {
		fmt.Println("Favorite song for ", k, ":")
		for _, v2 := range v {
			fmt.Print("\t", v2)
		}
		fmt.Println()
	}
}
