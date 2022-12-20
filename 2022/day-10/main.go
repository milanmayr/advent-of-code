package main

import (
	"strconv"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	print("The sum of the six signal strengths is: ")
	println(sumOfSixSignalStrengths())
}

func sumOfSixSignalStrengths() (X int) {
	input := utils.GetInput("example2-input")

	X = 1
	cycle := 1

	numberToAdd := 0
	for _, line := range input {
		if line == "noop" {
			cycle++
		}
		if line != "noop" {
			numberToAdd, _ = strconv.Atoi(line[5:])
			X += numberToAdd
			cycle ++
			cycle ++
		}
	}

	return X
}