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

// which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?
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