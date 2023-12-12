package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/milanmayr/advent-of-code/utils"
)

func main() {
	games := utils.GetInput("input")

	var sumOfPossibleGameIds int
	for _, line := range games {
		gameId, sets := parseLine(line)
		gamePossible := true
		for _, s := range sets {
			if !setPossible(s) {
				gamePossible = false
				break
			}
		}
		if gamePossible {
			sumOfPossibleGameIds += gameId
		}
	}
	fmt.Printf("The sum of possible Game IDs is %d\n", sumOfPossibleGameIds)

	// Part 2
	var sumOfSetPowers int
	for _, line := range games {
		_, sets := parseLine(line)
			minimumCubes := minimumCubes(sets)
			sumOfSetPowers += setPower(minimumCubes)
	}

	fmt.Printf("The sum of the power of minimum sets for each game is %d\n", sumOfSetPowers)
}

type colorSet struct {
	Red int
	Green int
	Blue int
}

func parseLine(line string) (gameId int, sets []colorSet) {
	// Get Game ID
	indexFirstSpace := strings.IndexByte(line, ' ')
	indexColon := strings.IndexByte(line, ':')
	gameId, err := strconv.Atoi(line[indexFirstSpace+1:indexColon])
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	
	// parse everything after "Game x:"
	colorSetsRawString := line[indexColon+2:]
	
	// split sets
	setsRawStrings := strings.Split(colorSetsRawString, ";")
	// handle each set
	for _, set := range setsRawStrings {
		colorSet := colorSet{}
		setTrimmedWhitespace := strings.TrimSpace(set)
		colorsRawStrings := strings.Split(setTrimmedWhitespace, ",")
		for _, v := range colorsRawStrings {
			valuesTrimmedWhitespace := strings.TrimSpace(v)
			values := strings.Split(valuesTrimmedWhitespace, " ")
			cubesCount, _ := strconv.Atoi(values[0])
			switch values[1] {
			case "red":
				colorSet.Red = cubesCount
			case "green":
				colorSet.Green = cubesCount
			case "blue":
				colorSet.Blue = cubesCount
			}
		}
		sets = append(sets, colorSet)
	}
	return 
}

// If the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes, would the set be possible?
func setPossible(set colorSet) bool {
	if set.Red > 12 {
		return false
	}
	if set.Green > 13 {
		return false
	}
	if set.Blue > 14 {
		return false
	}
	return true
}

// return minimum colorSet for a game
func minimumCubes(game []colorSet) (minimumCubes colorSet) {
	var red int
	var green int
	var blue int

	for _, set := range game {
		if set.Red > red {
			red = set.Red
		}
		if set.Green > green {
			green = set.Green
		}
		if set.Blue > blue {
			blue = set.Blue
		}
	}
	minimumCubes = colorSet{
		Red: red,
		Green: green,
		Blue: blue,
	}
	return minimumCubes
}

// return the power of a set
func setPower(set colorSet) int {
	return set.Blue * set.Red * set.Green
}