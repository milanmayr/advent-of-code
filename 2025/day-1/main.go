package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/milanmayr/advent-of-code/utils"
)

func main() {
	// fmt.Printf("Part 1: The password for the safe is %d\n", getPassword(false))
	fmt.Printf("Part 2: Using the 0x434C49434B password method, the password for the safe is %d\n", getPassword(true)) // 7883 is too high; 4778 not correct; 6946 not correct; 5849 not correct
}

type dial struct {
	position int
}

func (d *dial) rotateDial(clicks int, direction string) (timesZeroPassed int, err error) {
	timesZeroPassed = 0
	var newPosition int

	switch direction {
	case "R":
		newPosition = d.position + clicks%100
	case "L":
		newPosition = d.position - clicks%100
	default:
		return 0, errors.New("invalid direction. Must be 'R' or 'L'")
	}

	if newPosition < 0 {
		newPosition = 100 + newPosition
	} else if newPosition > 99 {
		newPosition = newPosition - 100
	}

	switch direction {
	case "R":
		timesZeroPassed = (d.position + clicks) / 100
	case "L":
		timesZeroPassed = utils.AbsoluteInteger(((d.position - clicks) / 100))
		// if (d.position - clicks ) < -100 {
		// 	timesZeroPassed = (d.position - clicks) / -100
		// } else if (d.position - clicks) < 0 {
		// 	timesZeroPassed = (d.position - clicks - 100) / -100
		// } else {
		// 	timesZeroPassed = (d.position - clicks) / 100
		// }

	}

	d.position = newPosition

	return timesZeroPassed, nil

}

func getPassword(useMethod0x434C49434B bool) (password int) {
	d := dial{
		position: 50,
	}

	passwordCounter := 0
	rotations := utils.GetInput("input")

	for _, rotation := range rotations {
		direction := rotation[0:1]
		clicks, err := strconv.Atoi(rotation[1:])
		if err != nil {
			fmt.Println("error converting input line to clicks int")
			os.Exit(1)
		}
		timesZeroPassed, err := d.rotateDial(clicks, direction)
		if err != nil {
			fmt.Printf("Error rotating dial for rotation %s", rotation)
			os.Exit(1)
		}

		if useMethod0x434C49434B {
			passwordCounter += timesZeroPassed
		} else {
			if d.position == 0 {
				passwordCounter++
			}
		}

	}

	return passwordCounter
}
