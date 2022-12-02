package utils

import (
	"bufio"
	"os"
	"strconv"
)

func stringToInteger(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func ReadInputTo2DArray() (array2d [][]int) {
	// This function reads an input text file and returns it as a 2D array
	// Each array inside the array is deliminated by empty lines in the input file

	// Initialize empty 2D array
	array2d = make([][]int, 1)

	// Read input file and populate the reindeer 2D array
	readFile, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	index := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			array2d[index] = append(array2d[index], stringToInteger(line))
		} else {
			index++
			array2d = append(array2d, make([]int, 0))
		}
	}

	readFile.Close()

	return array2d
}
