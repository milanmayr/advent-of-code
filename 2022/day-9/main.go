package main

import (
	"math"
	"strconv"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

// as long as delta of x - x and y - y is 1 or less, the head and tail are good
// if head is two steps up down left or right, tail moves one step in same direction
// if head and tail aren't touching and aren't in same row or column, tail moves one step diagonally to keep up
//

func main() {
	print("The tail of the rope visited ")
	print(positionsTailVisited())
	println(" positions at least once.")
}

type position struct {
	X int
	Y int
}

func positionsTailVisited() (positions int) {

	input := utils.GetInput("input")
	// initial position of head and tail is 0,0
	tail := &position{
		X: 0,
		Y: 0,
	}
	head := &position{
		X: 0,
		Y: 0,
	}

	positionsVisitedByTail := make(map[position]bool, 0)
	positionsVisitedByTail[*tail] = true

	for _, line := range input {
		direction := line[:1]
		distance, _ := strconv.Atoi(line[2:])

		for i := 1; i <= distance; i++ {
			// 1. move head
			switch direction {
			case "L":
				head.X--
			case "U":
				head.Y++
			case "R":
				head.X++
			case "D":
				head.Y--
			}
			
			// 2. move tail
			deltaX := int(math.Abs(float64(head.X-tail.X)))
			deltaY := int(math.Abs(float64(head.Y-tail.Y)))
			if deltaX > 1 || deltaY > 1 {
				if deltaY == 0 {
					tail.X = head.X - 1
				} else if deltaX == 0 {
					tail.Y = head.Y - 1
				} else if deltaY == 2 {
					tail.Y = head.Y - 1
					tail.X = head.X
				} else if deltaX == 2 {
					tail.X = head.X -1
					tail.Y = head.Y
				}
				positionsVisitedByTail[*tail] = true
			}}

	}
	return len(positionsVisitedByTail)
}