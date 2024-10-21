package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	r, err := regexp.Compile(`(\d+):`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

	/*
		bag_red := 12
		bag_green := 13
		bag_blue := 14
	*/

	example_game := "Game 1: 1 green, 6 red, 4 blue; 2 blue, 6 green, 7 red; 3 red, 4 blue, 6 green; 3 green; 3 blue, 2 green, 1 red"

	//id_matches := r.FindStringSubmatch(example_game)
	// id := id_matches[1]

	game_contents := strings.Split(example_game, ":")
	games := strings.Split(game_contents[1], ";")

	r, err = regexp.Compile(`(\d+) g`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

	for i := range len(games) {
		fmt.Println(games[i])
		matches := r.FindStringSubmatch(games[i])
		fmt.Println("There were ", matches[1], " green balls")
	}
}
