package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	r, err := regexp.Compile(`(\d+):`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

	sum_id := 0
	bag_red := 12
	bag_green := 13
	bag_blue := 14

	file, err := os.Open("day2-input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		id, _ := strconv.Atoi(r.FindStringSubmatch(line)[1])

		game_contents := strings.Split(line, ":")
		games := strings.Split(game_contents[1], ";")
		if validate_game(games, bag_red, bag_green, bag_blue) == true {
			sum_id += id
		}
	}

	fmt.Println("sum_id is: ", sum_id)
}

func validate_game(games []string, bag_red int, bag_green int, bag_blue int) bool {

	r_g, _ := regexp.Compile(`(\d+) g`)
	r_r, _ := regexp.Compile(`(\d+) r`)
	r_b, _ := regexp.Compile(`(\d+) b`)

	for i := range len(games) {
		num_green_balls := 0
		num_blue_balls := 0
		num_red_balls := 0

		matches := r_g.FindStringSubmatch(games[i])
		if len(matches) > 1 {
			num_green_balls, _ = strconv.Atoi(matches[1])
		}

		matches = r_b.FindStringSubmatch(games[i])
		if len(matches) > 1 {
			num_blue_balls, _ = strconv.Atoi(matches[1])
		}

		matches = r_r.FindStringSubmatch(games[i])
		if len(matches) > 1 {
			num_red_balls, _ = strconv.Atoi(matches[1])
		}

		if num_green_balls > bag_green || num_blue_balls > bag_blue || num_red_balls > bag_red {
			return false
		}
	}
	return true
}
