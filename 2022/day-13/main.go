package main

import (
	"strings"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	sumOfRightOrderedPairIndices()
}

func sumOfRightOrderedPairIndices() (sum int) {
	// 1-based indexing
	// for loop over every three lines -- first two lines are input, and skip blank third line
	// compare two lines
	// if the pair is in the right order, then add the index of the pair to the sum

	// for loop
	// each list becomes a slice
	//

	input := utils.GetInput("input")

	for i := 0; i < len(input)-1; i += 3 {
		packet1 := input[i]
		packet2 := input[i+1]

		var packet1lists []interface{}
		var packet2lists []interface{}

		packet1lists = createLists(packet1[1:])
		packet2lists = createLists(packet2[1:])

		print(packet1lists)
		print(packet2lists)

	}
	return sum
}

// func parseSlices(input string) (slices []interface{}) {
// 	for _, char := range input {
// 		if char == '[' {
// 			getList()
// 		}
// 	}
// 	return slices
// }

func createLists(str string) (list []interface{}) {
	// input string is the substring after the opening bracket
	var subStr string
	for index, char := range str {
		switch char {
		case '[':
			list = append(list, createLists(str[index+1:]))
		case ']':
			// convert string to list of integers
			if len(subStr) == 0 {
				list = append(list, make([]int, 0))
				return list
			} else {
				listStr := strings.Split(str[:index-1], ",")
				list = append(list, listStr)
				return list
			}
		default:
			subStr += string(char)
		}
	}

	return list
}
