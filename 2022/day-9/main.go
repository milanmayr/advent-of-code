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

	print("The tail of the rope with 10 knots visited ")
	print(positionsTailWith10KnotsVisited())
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
			deltaX := int(math.Abs(float64(head.X - tail.X)))
			deltaY := int(math.Abs(float64(head.Y - tail.Y)))
			if deltaX > 1 || deltaY > 1 {
				if deltaY == 0 {
					if head.X < tail.X {
						tail.X = head.X + 1
					} else {
						tail.X = head.X - 1
					}
				} else if deltaX == 0 {
					if head.Y < tail.Y {
						tail.Y = head.Y + 1
					} else {
						tail.Y = head.Y - 1
					}
				} else if deltaY == 2 {
					if head.Y < tail.Y {
						tail.Y = head.Y + 1
					} else {
						tail.Y = head.Y - 1
					}
					tail.X = head.X
				} else if deltaX == 2 {
					if head.X < tail.X {
						tail.X = head.X + 1
					} else {
						tail.X = head.X - 1
					}
					tail.Y = head.Y
				}
				positionsVisitedByTail[*tail] = true
			}
		}

	}
	return len(positionsVisitedByTail)
}

func positionsTailWith10KnotsVisited() (positions int) {
	input := utils.GetInput("input")
	// initial position of all knots is 0,0
	rope := make([]position, 10)
	for i := 0; i < 10; i++ {
		rope[i] = position{X: 0, Y: 0}
	}

	positionsVisitedByTail := make(map[position]bool, 0)
	positionsVisitedByTail[rope[9]] = true

	for _, line := range input {
		direction := line[:1]
		distance, _ := strconv.Atoi(line[2:])
		head := &rope[0]
		tail := &rope[9]

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

			// 2. move all knots behind head
			for i := 1; i <= 9; i++ {
				headOfCurrentKnot := rope[i-1]
				deltaX := int(math.Abs(float64(headOfCurrentKnot.X - tail.X)))
				deltaY := int(math.Abs(float64(headOfCurrentKnot.Y - tail.Y)))
				if deltaX > 1 || deltaY > 1 {
					if deltaY == 0 {
						if headOfCurrentKnot.X < tail.X {
							tail.X = headOfCurrentKnot.X + 1
						} else {
							tail.X = headOfCurrentKnot.X - 1
						}
					} else if deltaX == 0 {
						if headOfCurrentKnot.Y < tail.Y {
							tail.Y = headOfCurrentKnot.Y + 1
						} else {
							tail.Y = headOfCurrentKnot.Y - 1
						}
					} else if deltaY == 2 {
						if headOfCurrentKnot.Y < tail.Y {
							tail.Y = headOfCurrentKnot.Y + 1
						} else {
							tail.Y = headOfCurrentKnot.Y - 1
						}
						tail.X = headOfCurrentKnot.X
					} else if deltaX == 2 {
						if headOfCurrentKnot.X < tail.X {
							tail.X = headOfCurrentKnot.X + 1
						} else {
							tail.X = headOfCurrentKnot.X - 1
						}
						tail.Y = headOfCurrentKnot.Y
					}
					positionsVisitedByTail[*tail] = true
				}
			}
		}

	}
	return len(positionsVisitedByTail)
}
