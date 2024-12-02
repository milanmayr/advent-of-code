package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/milanmayr/advent-of-code/utils"
)

func main() {
	fmt.Printf("The total distance between the two lists of locations is %d\n", totalDistanceBetweenBothLists())
}

func totalDistanceBetweenBothLists() int {
	var totalDistance int
	input := utils.GetInput("input")

	// Create lists of locations
	var firstList []int
	var secondList []int

	// Import locations into maps created above and sort
	for _, row := range input {
		locations := strings.Split(row, "   ")
		firstNum, err := strconv.Atoi(locations[0])
		if err != nil {
			log.Fatalf("error converting %v to integer", firstNum)
		}
		secondNum, err := strconv.Atoi(locations[1])
		if err != nil {
			log.Fatalf("error converting %v to integer", secondNum)
		}

		firstList = append(firstList, firstNum)
		secondList = append(secondList, secondNum)

		sort.Ints(firstList)
		sort.Ints(secondList)
	}

	// import both lists into map
	allLocations := make(map[int]int, len(firstList))
	for i, location := range firstList {
		allLocations[location] = secondList[i]
	}

	// calculate total distance between all pairs of locations
	for firstLocation, secondLocation := range allLocations {
		totalDistance += absDiffInt(firstLocation, secondLocation)
	}
	return totalDistance
}

func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }
 