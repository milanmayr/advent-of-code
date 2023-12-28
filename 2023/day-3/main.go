package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

type schematicLocation struct {
	x int
	y int
}

type partNumberDigit struct {
	digit    rune
	location schematicLocation
}
type partNumber struct {
	number []partNumberDigit
}

func main() {
	var partNumbers []partNumber
	engineSchematic := utils.GetInput("input")
	for y, line := range engineSchematic {
		var partNumber partNumber
		for x := 0; x < len(line); x++ {
			if y == 127 && x == 139 {
				print("debug")
			}
			if unicode.IsDigit(rune(line[x])) {
				location := schematicLocation{
					x: x,
					y: y,
				}
				digit := partNumberDigit{
					digit:    rune(line[x]),
					location: location,
				}
				partNumber.number = append(partNumber.number, digit)
				if x == len(line)-1 && partNumber.number != nil {
					partNumbers = append(partNumbers, partNumber)
				}
			} else {
				if partNumber.number != nil {
					partNumbers = append(partNumbers, partNumber)
				}
				partNumber.number = nil
			}
		}
	}

	var sum int
	for _, num := range partNumbers {
		// check if each part number is adjacent to a symbol
		if partNumberAdjacentToSymbol(num, engineSchematic) {
			var stringNum string
			for _, d := range num.number {
				stringNum += string(d.digit)
			}
			intNum, err := strconv.Atoi(stringNum)
			if err != nil {
				print(err.Error())
				os.Exit(1)
			}
			sum += intNum
		}
	}

	fmt.Printf("The sum of all engine schematic part numbers is %d\n", sum)
}

func partNumberAdjacentToSymbol(partNum partNumber, engineSchematic []string) bool {
	var locations []schematicLocation
	num := partNum.number
	for i := 0; i < len(num); i++ {
		currentDigitLocation := num[i].location

		// for all digits, check above location
		if currentDigitLocation.y > 0 {
			locations = append(locations, schematicLocation{
				x: currentDigitLocation.x,
				y: currentDigitLocation.y - 1,
			})
		}
		// for all digits, check below location
		if currentDigitLocation.y < len(engineSchematic)-1 {
			locations = append(locations, schematicLocation{
				x: currentDigitLocation.x,
				y: currentDigitLocation.y + 1,
			})
		}

		// for first digits, also check to the left of, both left diagonals
		if i == 0 {
			// check left of digit
			if currentDigitLocation.x > 0 {
				locations = append(locations, schematicLocation{
					x: currentDigitLocation.x - 1,
					y: currentDigitLocation.y,
				})
			}

			// check top left diagonal
			if currentDigitLocation.x > 0 && currentDigitLocation.y > 0 {
				locations = append(locations, schematicLocation{
					x: currentDigitLocation.x - 1,
					y: currentDigitLocation.y - 1,
				})
			}

			// check bottom left diagonal
			if currentDigitLocation.x > 0 && currentDigitLocation.y < len(engineSchematic)-1 {
				locations = append(locations, schematicLocation{
					x: currentDigitLocation.x - 1,
					y: currentDigitLocation.y + 1,
				})
			}
		}

		// for last digits, also check to the right of, both right diagonals
		if i == len(partNum.number)-1 {
			// check right of digit
			if currentDigitLocation.x < len(engineSchematic[currentDigitLocation.y])-1 {
				locations = append(locations, schematicLocation{
					x: currentDigitLocation.x + 1,
					y: currentDigitLocation.y,
				})
			}

			// check top right diagonal
			if currentDigitLocation.x < len(engineSchematic[currentDigitLocation.y])-1 && currentDigitLocation.y > 0 {
				locations = append(locations, schematicLocation{
					x: currentDigitLocation.x + 1,
					y: currentDigitLocation.y - 1,
				})
			}

			// check bottom right diagonal
			if currentDigitLocation.x < len(engineSchematic[currentDigitLocation.y])-1 && currentDigitLocation.y < len(engineSchematic)-1 {
				locations = append(locations, schematicLocation{
					x: currentDigitLocation.x + 1,
					y: currentDigitLocation.y + 1,
				})
			}
		}
	}

	for _, l := range locations {
		r := rune(engineSchematic[l.y][l.x])
		if !unicode.IsDigit(r) {
			if r != '.' {
				return true
			}
		}
	}

	return false
}
