package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/milanmayr/advent-of-code/utils"
)

func main() {
	fmt.Printf("The sum of all calibration values is %d\n", sumOfAllCalibrationValues())
	fmt.Printf("The sum of all calibration values, including word numbers, is %d", sumOfAllCalibrationValuesIncludingStringNumbers())
}

func sumOfAllCalibrationValues() (sumOfAllCalibrationValues int) {
	calibrationDocument := utils.GetInput("input")

	// for each line, get first and last digit in string
	for _, line := range calibrationDocument {
		// get first and last digit in string
		var digits []rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}
		sumOfAllCalibrationValues += convertRuneSliceToInt(digits)
	}
	return sumOfAllCalibrationValues
}

var wordNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func sumOfAllCalibrationValuesIncludingStringNumbers() (sumOfAllCalibrationValues int) {
	calibrationDocument := utils.GetInput("input")

	for _, line := range calibrationDocument {
		number := firstTwoDigitsInStringIncludingNumberWords(line)
		sumOfAllCalibrationValues += number
	}
	return sumOfAllCalibrationValues
}

func firstTwoDigitsInStringIncludingNumberWords(line string) (number int) {
	var digits []rune

	for i := 0; i <= len(line)-1; i++ {
		if unicode.IsDigit(rune(line[i])) {
			digits = append(digits, rune(line[i]))
			continue
		}
		// substring exists
		for j := i + 1; j <= len(line)-1; j++ {
			if unicode.IsDigit(rune(line[j])) {
				digits = append(digits, rune(line[j]))
				i = j
				break
			}
			restOfString := line[i : j+1]
			for k := range wordNumbers {
				if strings.Contains(restOfString, k) {
					char := strconv.Itoa(wordNumbers[k])
					digits = append(digits, rune(char[0]))
					i = j
					break
				}
			}
		}
	}
	return convertRuneSliceToInt(digits)
}

func convertRuneSliceToInt(digits []rune) (number int) {
	// convert first and last digit to two-digit number
	number, err := strconv.Atoi(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]))
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(10)
	}
	return number
}