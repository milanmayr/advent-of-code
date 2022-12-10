package main

import (
	"strconv"
	"strings"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	print("The crates on top of each stack are: ")
	println(topCrates())
	print("CrateMover 9001 -- The crates on top of each stack are: ")
	println(topCratesPartTwo())
}

func topCrates() (topCrates string) {
	// get first 9 lines and split into 9 queues
	// create the moving algorithm
	input := utils.GetInput("input")

	// Create set of arrays representing stacks
	// Stack 0 will be empty to make index association easy
	stacks := getStacks()

	// Crate moving logic
	for line := 10; line <= len(input)-1; line++ {
		numberOfCrates, _ := strconv.Atoi(strings.Split(input[line], " ")[1])
		source, _ := strconv.Atoi(strings.Split(input[line], " ")[3])
		destination, _ := strconv.Atoi(strings.Split(input[line], " ")[5])

		for i := 1; i <= numberOfCrates; i++ {
			// get crate
			crate := stacks[source][len(stacks[source])-1]
			// remove crate from source
			stacks[source] = stacks[source][:len(stacks[source])-1]
			// add crate to destination
			stacks[destination] = append(stacks[destination], crate)
		}
	}

	for stack := 1; stack < len(stacks); stack++ {
		topCrates = topCrates + stacks[stack][len(stacks[stack])-1]
	}

	return topCrates
}

func topCratesPartTwo() (topCrates string) {
	input := utils.GetInput("input")
	stacks := getStacks()

	// Crate moving logic
	for line := 10; line <= len(input)-1; line++ {
		numberOfCrates, _ := strconv.Atoi(strings.Split(input[line], " ")[1])
		source, _ := strconv.Atoi(strings.Split(input[line], " ")[3])
		destination, _ := strconv.Atoi(strings.Split(input[line], " ")[5])

		// get top n crates
		cratesToMove := stacks[source][len(stacks[source])-numberOfCrates:len(stacks[source])]
		// remove top n crates from source
		stacks[source] = stacks[source][:len(stacks[source])-numberOfCrates]
		// add top n crates to destination in order
		stacks[destination] = append(stacks[destination], cratesToMove...)
	}

	for stack := 1; stack < len(stacks); stack++ {
		topCrates = topCrates + stacks[stack][len(stacks[stack])-1]
	}

	return topCrates
}

func getStacks() (stacks [][]string) {
	// get first 9 lines and split into 9 queues
	// create the moving algorithm
	input := utils.GetInput("input")

	// Create set of arrays representing stacks
	// Stack 0 will be empty to make index association easy
	stacks = make([][]string, 10)

	for line := 7; line >= 0; line-- {
		stack := 1
		for charPosition := 1; charPosition < 35; charPosition += 4 {
			crate := strings.Split(input[line], "")[charPosition]
			if (crate != " ") {
				stacks[stack] = append(stacks[stack], crate)
			}
			stack++
		}
	}

	return stacks
}