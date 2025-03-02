package utils

import (
	"bufio"
	"errors"
	"os"
)

func GetInput(inputFile string) (challengeInput []string) {
	var input []string

	readFile, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		input = append(input, line)
	}

	readFile.Close()

	return input
}

func IndexOf(element string, data []string) (int, error) {
	for k, v := range data {
		if element == v {
			return k, nil
		}
	}
	return 0, errors.New("element not found in string array") // value not found
}

func AbsoluteDifferenceInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
