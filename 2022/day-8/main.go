package main

import (
	"strconv"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	print("The number of trees visible from the edges is: ")
	println(treesVisibleFromEdge())
}

func treesVisibleFromEdge() (visibleTrees int) {
	input := utils.GetInput("input")

	trees := make([][]int, 99)
	// Create 2D array from input
	for row, line := range input {
		for _, char := range line {
			tree, _ := strconv.Atoi(string(char))
			trees[row] = append(trees[row], tree)
		}
	}

	// calculate number of trees on edge; these will automatically be visible
	visibleTrees = 2*len(trees[1]) + 2*len(trees) - 4

	// for loop iterating over all trees
	for row := 1; row < len(trees)-1; row++ {
		for column := 1; column < len(trees[0])-1; column++ {
			tree := trees[row][column]
			visibleLeft := true
			visibleAbove := true
			visibleRight := true
			visibleBelow := true
			// check all trees to the left
			for left := column - 1; left >= 0; left-- {
				if tree <= trees[row][left] {
					visibleLeft = false
					break
				}
			}
			// check all trees above
			for above := row - 1; above >= 0; above-- {
				if tree <= trees[above][column] {
					visibleAbove = false
					break
				}
			}
			// check all trees to the right
			for right := column + 1; right < len(trees[0]); right++ {
				if tree <= trees[row][right] {
					visibleRight = false
					break
				}
			}
			// check all trees below
			for below := row + 1; below < len(trees); below++ {
				if tree <= trees[below][column] {
					visibleBelow = false
					break
				}
			}

			if visibleLeft || visibleRight || visibleAbove || visibleBelow {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}
