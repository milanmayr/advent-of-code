package main

import (
	"strings"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	print("The sum of priorities for all the shared items is: ")
	println(sumPriorities())
	print("The sum of priorities for the shared badges of each three-elf group is: ")
	println(sumPrioritiesPartTwo())
}

var PriorityAlphabet string = "0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func sumPriorities() (sumOfPriorities int) {
	// var sumOfPriorities int
	input := utils.GetInput("input")

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
					priorityValue, err := utils.IndexOf(commonChar, strings.Split(PriorityAlphabet, ""))
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

func sumPrioritiesPartTwo() (sumPriorities int) {
	input := utils.GetInput("input")

	sumPriorities = 0
	for rucksack := 0; rucksack < len(input); rucksack += 3 {
		firstRucksack := input[rucksack]
		secondRucksack := input[rucksack+1]
		thirdRucksack := input[rucksack+2]

		firstChars := []rune(firstRucksack)
		secondChars := []rune(secondRucksack)
		thirdChars := []rune(thirdRucksack)

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
					for k := 0; k < len(thirdChars); k++ {
						if commonCharFound {
							break
						}
						if string(secondChars[j]) == string(thirdChars[k]) {
							commonChar := string(firstChars[i])
							priorityValue, err := utils.IndexOf(commonChar, strings.Split(PriorityAlphabet, ""))
							if err != nil {
								panic("element was not found in string array")
							}
							sumPriorities += priorityValue
							commonCharFound = true
							break
						}
					}
				}
			}
			
		}
	
	}
	return sumPriorities
}
