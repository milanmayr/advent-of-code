package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/milanmayr/advent-of-code/utils"
)

type listsOfReports [][]int

func main() {
	fmt.Printf("The count of safe reports is %d\n", countSafeReports())
	fmt.Printf("The count of safe reports with the Problem Dampener is %d\n", countSafeReportsWithProblemDampener())
}

// Parse input
func getReports() (lol listsOfReports) {
	input := utils.GetInput("input")

	for _, row := range input {
		var report []int
		stringReport := strings.Split(row, " ")
		for _, stringLevel := range stringReport {
			level, err := strconv.Atoi(stringLevel)
			if err == nil {
				report = append(report, level)
			} else {
				print(err)
				os.Exit(1)
			}
		}
		lol = append(lol, report)
	}

	return lol
}

// Determine if a report is safe
func reportIsSafe(report []int) bool {
	var ascending int
	var descending int

	for i, level := range report {
		if i == len(report)-1 {
			break
		}

		nextLevel := report[i+1]
		// Because numbers must be either ascending or descending, equal numbers should fail and return false
		if level == nextLevel {
			return false
		}

		// Because the difference between two numbers must be at least one and less than three, any absolute
		// differences larger than 3 should fail and return false

		if utils.AbsoluteDifferenceInt(level, nextLevel) > 3 {
			return false
		}

		if level > nextLevel {
			descending++
		} else if level < nextLevel {
			ascending++
		}
	}

	// If the ascending and descending counts are both above 0, return false
	if ascending > 0 && descending > 0 {
		return false
	}

	return true
}

func countSafeReports() int {
	var safeReports int
	reports := getReports()

	for _, report := range reports {
		if reportIsSafe(report) {
			safeReports++
		}
	}

	return safeReports
}

func countSafeReportsWithProblemDampener() int {
	var safeReports int
	reports := getReports()

	for _, report := range reports {
		if reportIsSafe(report) {
			safeReports++
		} else {
			for i := range report {
				unsafeReport := make([]int, len(report))
				copy(unsafeReport, report)
				reportWithoutLevel := append(unsafeReport[:i], unsafeReport[i+1:]...)
				if reportIsSafe(reportWithoutLevel) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}