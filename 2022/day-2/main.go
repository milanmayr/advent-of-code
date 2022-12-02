package main

import (
	"bufio"
	"os"
)

func main() {
	print("The total score would be ")
	println(totalScore())
	print("The total score for the second part would be ")
	println(totalScorePartTwo())
}

func totalScore() (totalScore int) {
	readFile, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		totalScore += rockPaperScissors(string(line[2]), string(line[0]))
	}

	readFile.Close()

	return totalScore
}

func totalScorePartTwo() (totalScore int) {
	readFile, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		opponent := string(line[0])
		outcome := string(line[2])
		var me string

		switch outcome {
		case "X":
			// lose
			switch opponent {
			case "A":
				me = "Z"
			case "B":
				me = "X"
			case "C":
				me = "Y"
			}
		case "Y":
			// draw
			switch opponent {
			case "A":
				me = "X"
			case "B":
				me = "Y"
			case "C":
				me = "Z"
			}
		case "Z":
			// win
			switch opponent {
			case "A":
				me = "Y"
			case "B":
				me = "Z"
			case "C":
				me = "X"
			}
		}
		totalScore += rockPaperScissors(me, opponent)
	}

	readFile.Close()

	return totalScore
}

func rockPaperScissors(me string, opponent string) (score int) {
	// rock loses to paper, wins against scissors
	// paper loses to scissors, wins against rock
	// scissors loses to rock, wins against paper
	score = 0

	switch me {
	case "X":
		score += 1
		if opponent == "C" {
			score += 6
		}
		if opponent == "A" {
			score += 3
		}
	case "Y":
		score += 2
		if opponent == "A" {
			score += 6
		}
		if opponent == "B" {
			score += 3
		}
	case "Z":
		score += 3
		if opponent == "B" {
			score += 6
		}
		if opponent == "C" {
			score += 3
		}
	}

	return score
}
