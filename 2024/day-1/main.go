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
	fmt.Printf("The similarity score for the two lists of locations is %d\n", similarityScoreOfBothLists())
}

type listsOfLocationIDs struct {
	firstList  []int
	secondList []int
}

func parseLocationIDLists() listsOfLocationIDs {
	var lists listsOfLocationIDs
	input := utils.GetInput("input")

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

		lists.firstList = append(lists.firstList, firstNum)
		lists.secondList = append(lists.secondList, secondNum)
	}

	return lists

}

func totalDistanceBetweenBothLists() int {
	var totalDistance int

	lists := parseLocationIDLists()

	sort.Ints(lists.firstList)
	sort.Ints(lists.secondList)

	// import both lists into map
	allLocations := make(map[int]int, len(lists.firstList))
	for i, location := range lists.firstList {
		allLocations[location] = lists.secondList[i]
	}

	// calculate total distance between all pairs of locations
	for firstLocation, secondLocation := range allLocations {
		totalDistance += absDiffInt(firstLocation, secondLocation)
	}
	return totalDistance
}

// sum of all the numbers in the first (left) list after each number has been multiplied by the number of times it appears in the right list
func similarityScoreOfBothLists() int {
	var similarityScore int

	lists := parseLocationIDLists()

	// create map of how many times each number in the first list appears in the right list
	occurrenceMap := calculateOccurrences(lists)

	for _, v := range lists.firstList {
		similarityScore += v * occurrenceMap[v]
	}

	return similarityScore
}

type numberOccurrences map[int]int

func calculateOccurrences(lists listsOfLocationIDs) numberOccurrences {
	var occurrenceMap = make(map[int]int)

	sort.Ints(lists.firstList)
	sort.Ints(lists.secondList)

	fl := lists.firstList
	sl := lists.secondList

	// // remove duplicates from first list 
	// for i, v := range fl {
	// 	if i == len(fl) - 1 {break}
	// 	if v == fl[i + 1] {
	// 		fl = append(fl[0:v], fl[i+1:len(fl)-1]...)
	// 	}
	// }

	for _, v := range fl {
		occurrenceMap[v] = 0
	}

	// create map based on a sorted second list.
	for _, firstValue := range fl {
		for _, secondValue := range sl {
			if firstValue == secondValue {
				occurrenceMap[firstValue]++
			}
		}
	}
	return occurrenceMap
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

