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
)

func main() {
	calibration_value_sum := 0

	file, err := os.Open("day1-input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line_no := 0
	for scanner.Scan() {
		line := scanner.Text()
		calibration_value_sum += get_digit(line)
		line_no += 1
	}

	fmt.Println(calibration_value_sum)
}

func get_digit(line string) int {
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
