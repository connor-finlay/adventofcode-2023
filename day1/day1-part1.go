/*
On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	calibration_value_sum := 0

	file, err := os.Open("example-input-pt2")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line_no := 0
	for scanner.Scan() {
		line := scanner.Text()
		calibration_value_sum += get_digit_pt1(line)
		part2(line)
		line_no += 1
	}

	fmt.Println("Part 1: ", calibration_value_sum)
}

func get_digit_pt1(line string) int {
	digits := []byte{0x30, 0x30}

	for i := 0; i < len(line); i++ {
		if line[i] >= 0x30 && line[i] <= 0x39 {
			digits[0] = line[i]
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= 0x30 && line[i] <= 0x39 {
			digits[1] = line[i]
			break
		}
	}

	calibration_str := string(digits)
	calibration_value, err := strconv.Atoi(calibration_str)
	if err != nil {
		fmt.Println(err)
	}
	return calibration_value
}

func part2(line string) {

	/*
		Check if there is a digit like '0' or '7'
		Take note of the index, if the index greater than 2 (three letter minimum to spell digit)
		Look for spelt digit

		Do the same for the last digit
	*/

	first_digit := -1

	// get the first_digit
	index, first_digit := get_first_digit(line)
	if first_digit == -1 {
		index, first_digit = get_digit_string(line)
	} else if index > 2 { // minimum length of digit is 3 letters
		_, digit_index := get_digit_string(line[0:index])
		if digit_index > -1 {
			first_digit = digit_index
		}
	}
	fmt.Println("first digit ", first_digit)
}

func get_first_digit(line string) (int, int) {
	for i := range len(line) {
		if line[i] >= 0x30 && line[i] <= 0x39 {
			return i, int(line[i] - 0x30)
		}
	}
	return -1, -1
}

func get_digit_string(line string) (int, int) {
	index := -1
	digit := -1
	for i := range len(digits) {
		sub_index := strings.Index(line, digits[i])
		if sub_index != -1 && (sub_index <= index || index == -1) {
			index = sub_index
			digit = i + 1
		}
	}
	return index, digit
}
