package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/milanmayr/advent-of-code/utils"
)

func main() {
	fmt.Printf("The sum of all calibration values is %d", sumOfAllCalibrationValues())
}

func sumOfAllCalibrationValues() (sumOfAllCalibrationValues int) {
	calibrationDocument := utils.GetInput("input")

	// for each line, get first and last digit in string
	for _, line := range calibrationDocument {
		// get first and last digit in string
		var digits []rune
		for _, char := range(line) {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}

		// convert first and last digit to two-digit number
		number, err := strconv.Atoi(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]))
		if err != nil {
			log.Fatalf(err.Error())
		}

		// add number to calibration sum
		sumOfAllCalibrationValues += number
	}

	return sumOfAllCalibrationValues
}