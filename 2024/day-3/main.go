package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/milanmayr/advent-of-code/utils"
)

func main() {
    fmt.Printf("The sum of all valid multiplier functions is %d\n", sumOfAllValidMultiplications())
}

func sumOfAllValidMultiplications() int {
	input := utils.GetInput("input")
	var sum int

	multiplierFunctionRegex, err := regexp.Compile(`mul\(\d+,\d+\)`)
	factorRegex, err := regexp.Compile(`\d+`)
	if err != nil {
		print(err)
		os.Exit(1)
	}

	for _, row := range input {
		validMultiplierFunctions := multiplierFunctionRegex.FindAllString(row, -1)

		for _, function := range validMultiplierFunctions {
			factors := factorRegex.FindAllString(function, -1)
			if len(factors) != 2 {
				fmt.Printf("More than 2 multiplier found in function `%s`. Exiting", function)
				os.Exit(1)
			}
			factor1, err := strconv.Atoi(factors[0])
			if err != nil {
				fmt.Printf("Error converting factor to int for string-typed factor `%s`. Exiting", factors[0])
			}
			factor2, err := strconv.Atoi(factors[1])
			if err != nil {
				fmt.Printf("Error converting factor to int for string-typed factor `%s`. Exiting", factors[1])
			}
			sum += factor1 * factor2
		}
	}
	
	return sum
}