package main

import (
	"strconv"
	"strings"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	print("The total of assignment pairs in which one fully contains the other is: ")
	println(countContainedPairs())
	print("The total of assignment pairs that overlap is: ")
	println(countOverlappingPairs())
}

func countContainedPairs() (count int) {
	input := utils.GetInput("input")
	count = 0
	for _, line := range input {
		pair := strings.Split(line, ",")
		first := strings.Split(pair[0], "-")
		second := strings.Split(pair[1], "-")

		firstMin, _ := strconv.Atoi(first[0])
		firstMax, _ := strconv.Atoi(first[1])
		secondMin, _ := strconv.Atoi(second[0])
		secondMax, _ := strconv.Atoi(second[1])

		if firstMin <= secondMin && secondMax <= firstMax {
			count++
		} else if secondMin <= firstMin && firstMax <= secondMax {
			count++
		} else {
			continue
		}
	}

	return count
}

func countOverlappingPairs() (count int) {
	count = 0

	input := utils.GetInput("input")
	count = 0
	for _, line := range input {
		pair := strings.Split(line, ",")
		first := strings.Split(pair[0], "-")
		second := strings.Split(pair[1], "-")

		firstMin, _ := strconv.Atoi(first[0])
		firstMax, _ := strconv.Atoi(first[1])
		secondMin, _ := strconv.Atoi(second[0])
		secondMax, _ := strconv.Atoi(second[1])

		if (secondMin <= firstMax && secondMin >= firstMin) || (firstMin <= secondMax && firstMin >= secondMin) {
			count++
		}
	}

	return count
}