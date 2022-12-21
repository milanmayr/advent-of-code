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

	partTwo()
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

func partTwo() {
	input := utils.GetInput("input")

	// create screen
	screen := make([][]rune, 6)
	for line := 0; line < 6; line++ {
		for i := 0; i <= 39; i++ {
			screen[line] = append(screen[line], '.')
		}
	}

	// handle cycles
	cycle := 0
	X := 1
	spritePosition := []int{0, 1, 2}

	for _, line := range input {
		if line == "noop" {
			cycleTickAndDrawScreen(&cycle, &X, &screen, &spritePosition)
		} else if line != "noop" {
			numberToAdd, _ := strconv.Atoi(line[5:])
			cycleTickAndDrawScreen(&cycle, &X, &screen, &spritePosition)
			cycleTickAndDrawScreen(&cycle, &X, &screen, &spritePosition)
			X += numberToAdd
			spritePosition = []int{X - 1, X, X + 1}
		}
	}
	// print screen
	for _, line := range screen {
		for _, pixel := range line {
			print(string(pixel))
		}
		println()
	}
}

func cycleTickAndDrawScreen(cycle *int, X *int, screen *[][]rune, spritePosition *[]int) {
	*cycle++

	row := *cycle / 40
	pixel := *cycle%40 - 1
	if *cycle % 40 == 0 {
		row = *cycle / 40 -1
		pixel = 39
	}
	// draw screen
	for _, position := range *spritePosition {
		if position == pixel {
			(*screen)[row][pixel] = '#'
		}
	}
}
