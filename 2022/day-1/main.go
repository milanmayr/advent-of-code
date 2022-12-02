package main

import (
	"sort"

	"github.com/milanmayr/advent-of-code/2022/day-1/utils"
)

func main() {
	print("The highest calories that one reindeer is carrying is ")
	println(highestCalories())
	print("The sum of the reindeer's top 3 total calories is ")
	println(totalOfThreeHighestCalories())
}

func highestCalories() (highestCalories int) {

	// Initialize empty 2D array [reindeer][calories]
	reindeer := utils.ReadInputTo2DArray()

	// Calculate highest calories per reindeer
	highestCalories = 0
	for index := range reindeer {
		reindeersCalories := 0
		for _, calories := range reindeer[index] {
			reindeersCalories += calories
		}
		if reindeersCalories > highestCalories {
			highestCalories = reindeersCalories
		}
	}

	return highestCalories
}

func totalOfThreeHighestCalories() (totalOfThreeHighestCalories int) {

	totalOfThreeHighestCalories = 0

	// Initialize empty 2D array [reindeer][calories]
	reindeer := utils.ReadInputTo2DArray()

	reindeerTotals := make([]int, 0)
	// Calculate three highest calories per reindeer
	for index := range reindeer {
		reindeersCalories := 0
		for _, calories := range reindeer[index] {
			reindeersCalories += calories
		}
		reindeerTotals = append(reindeerTotals, reindeersCalories)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(reindeerTotals)))

	totalOfThreeHighestCalories = reindeerTotals[0] + reindeerTotals[1] + reindeerTotals[2]

	return totalOfThreeHighestCalories
}
