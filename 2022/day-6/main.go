package main

import (
	"strings"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	print("The first start-of-packet marker ends at position: ")
	println(firstStartOfPacketMarkerPosition())
	print("The first start-of-message marker ends at position: ")
	println(firstStartOfMessageMarkerPosition())
}

func firstStartOfPacketMarkerPosition() (position int) {
	input := utils.GetInput("input")
	// input only has one line
	datastream := strings.Split(input[0], "")

	// start of a packet is indicated by four characters that are different
	// get position of character that ends the first start-of-packet marker

	potentialMarker := datastream[0:4]
	
	for char := 4; char < len(datastream)-1; char++ {
		if !isMarkerUnique(potentialMarker) {
			potentialMarker = potentialMarker[1:]
			potentialMarker = append(potentialMarker, datastream[char])
		} else {
			position = char
			break
		}
	}

	return position
}

func firstStartOfMessageMarkerPosition() (position int) {
	input := utils.GetInput("input")
	// input only has one line
	datastream := strings.Split(input[0], "")

	// start of a packet is indicated by four characters that are different
	// get position of character that ends the first start-of-packet marker

	potentialMarker := datastream[0:14]
	
	for char := 14; char < len(datastream)-1; char++ {
		if !isMarkerUnique(potentialMarker) {
			potentialMarker = potentialMarker[1:]
			potentialMarker = append(potentialMarker, datastream[char])
		} else {
			position = char
			break
		}
	}

	return position
}

func isMarkerUnique(marker []string) (bool) {
	for char := 0; char < len(marker)-1; char++ {
		for otherchar := char+1; otherchar < len(marker); otherchar++ {
			if marker[char] == marker[otherchar] {
				return false
			}
		}
	}
	return true
}