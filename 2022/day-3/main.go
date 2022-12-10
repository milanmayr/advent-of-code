package main

import (
	"strings"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	print("The sum of priorities for all the shared items is: ")
	print(sumPriorities())
}

func sumPriorities() (sumOfPriorities int) {
	// var sumOfPriorities int
	input := utils.GetInput("input")

	priorityAlphabet := "0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// split each line into two arrays. Length of each is half of the whole line
	// compare both arrays to find the common character
	// add priority of shared item to sumOfPriorities

	sumOfPriorities = 0
	for _, rucksack := range input {

		stringLength := len(rucksack)
		firstCompartment := rucksack[:stringLength/2]
		secondCompartment := rucksack[stringLength/2:]

		firstChars := []rune(firstCompartment)
		secondChars := []rune(secondCompartment)

		var commonCharFound bool
		for i := 0; i < len(firstChars); i++ {
			if commonCharFound {
				break
			}
			for j := 0; j < len(secondChars); j++ {
				if commonCharFound {
					break
				}
				if string(firstChars[i]) == string(secondChars[j]) {
					commonChar := string(firstChars[i])
					priorityValue, err := utils.IndexOf(commonChar, strings.Split(priorityAlphabet, ""))
					if err != nil {
						panic("element was not found in string array")
					}
					sumOfPriorities += priorityValue
					commonCharFound = true
					break
				}
			}
		}
	}

	return sumOfPriorities
}
