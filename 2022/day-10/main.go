package main

import (
	"fmt"
	"strconv"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	sum := sumOfSixSignalStrengths()
	print("The sum of all six signal strengths is ")
	println(sum)
}

func sumOfSixSignalStrengths() (sum int) {
	input := utils.GetInput("input")

	cycle := 0
	X := 1
	thCycles := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}

	for _, line := range input {
		if line == "noop" {
			cycleTick(&cycle, thCycles, &X, &sum)
		} else if line != "noop" {
			numberToAdd, _ := strconv.Atoi(line[5:])
			cycleTick(&cycle, thCycles, &X, &sum)
			cycleTick(&cycle, thCycles, &X, &sum)
			X += numberToAdd
		}
	}
	return sum
}

func cycleTick(cycle *int, thCycles map[int]bool, X *int, sum *int) {
	msg := "The signal strength during the %dth cycle is %d"

	*cycle++

	if thCycles[*cycle] {
		signalStrength := *X * *cycle
		fmt.Printf(msg, *cycle, signalStrength)
		println()
		*sum += signalStrength
	}

}
