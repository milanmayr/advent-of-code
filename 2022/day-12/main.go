package main

import (
	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {

	print("The fewest number of steps to get the best signal is: ")
	println(fewestStepsToBestSignal())
	// cannot move to a square that's two levels higher
	// can move to a square that's any number of levels lower
	// can use rune Unicode values to decide priority

}

type position struct {
	value rune
	moveUp *position
	moveDown *position
	moveRight *position
	moveLeft *position
}

func fewestStepsToBestSignal() (steps int) {
	input := utils.GetInput("input")

	// Create heightmap using struct
	heightmap := make([][]position, 0)
	for y := 0; y < 41; y++ {
		runes := []rune(input[y])
		// add y line
		heightmap = append(heightmap, make([]position, 0))
		for x := 0; x < len(runes)-1; x++ {
			heightmap[y] = append(heightmap[y], position{
				value: runes[x],
			})
		}
	}

	// // Create []rune heightmap
	// heightmap := make([][]rune, 0)
	// for y := 0; y < 41; y++ {
	// 	runes := []rune(input[y])
	// 	heightmap = append(heightmap, runes)
	// }

	// find positions of S and E
	var startingPoint *position
	var endPoint *position
	for _, line := range heightmap {
		for _, pos := range line {
			if pos.value == 83 {
				startingPoint = &pos
			} else if pos.value  == 69 {
				endPoint = &pos
			}
		}
	}

	print(startingPoint)
	print(endPoint)

	// Populate heightmap with possible moves
	for y := 0; y < 41; y++ {
		for x := 0; x < 79; x++ {
			pos := &heightmap[y][x]
			if y > 0 && heightmap[y-1][x].value <= pos.value+1 {
				pos.moveUp = &heightmap[y-1][x]
			}
			// can move down
			if y < 40 && heightmap[y+1][x].value <= pos.value+1 {
				pos.moveDown = &heightmap[y+1][x]
			}
			// can move left
			if x > 0 && heightmap[y][x-1].value <= pos.value+1 {
				pos.moveLeft = &heightmap[y][x-1]
			}
			// can move right
			if x < 78 && heightmap[y][x+1].value <= pos.value+1 {
				pos.moveRight = &heightmap[y][x+1]
			}
		}
	}
	for y, line := range heightmap {
		// var upPosition *position
		// var downPosition *position
		// var rightPosition *position
		// var leftPosition *position

		for x, pos := range line {
			// can move up
			if y > 0 && heightmap[y-1][x].value <= pos.value+1 {
				pos.moveUp = &heightmap[y-1][x]
			}
			// can move down
			if y < 40 && heightmap[y+1][x].value <= pos.value+1 {
				pos.moveDown = &heightmap[y+1][x]
			}
			// can move left
			if x > 0 && heightmap[y][x-1].value <= pos.value+1 {
				pos.moveLeft = &heightmap[y][x-1]
			}
			// can move right
			if x < 78 && heightmap[y][x+1].value <= pos.value+1 {
				pos.moveRight = &heightmap[y][x+1]
			}
		}
	}

	// Get shortest number of steps
	// steps = calculateFewestSteps(heightmap, *startingPoint, *endPoint)

	return steps
}

func canMove(current rune, next rune) bool {
	// Based on rune values of characters

	// if current is S (start point) and next is a, allow
	// else if next is E (end point), allow
	// else if next is 1 above current or equal to current or below current, allow
	// otherwise, deny
	if (current == 83) && (next == 93) {
		return true
	} else if next == 69 {
		return true
	} else if (current+1 == next) || (next <= current) {
		return true
	} else {
		return false
	}
}

func possibleMoves(currentPosition *position) (positions []position) {
	// above := &
	return positions
}

func calculateFewestSteps(heightmap [][]position, start *position, end *position) (steps int) {
	
	return steps
}
