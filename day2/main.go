package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, err := regexp.Compile(`(\d+):`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

	id_matches := r.FindStringSubmatch("Game 1: 1 green, 6 red, 4 blue; 2 blue, 6 green, 7 red; 3 red, 4 blue, 6 green; 3 green; 3 blue, 2 green, 1 red")

	id := id_matches[1]

}
