package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	print("After 20 rounds of stuff-slinging simian shenanigans, the level of monkey business is: ")
	println(getMonkeyBusinessLevel(20, true))
	print("After 10,000 rounds of stuff-slinging simian shenanigans and no longer dividing the worry level by 3, the level of monkey business is: ")
	println(getMonkeyBusinessLevel(10000, false))
}

type monkey struct {
	StartingItems []int64
	Operation     struct {
		Operator string
		Number   int64
	}
	DivisibleBy         int64
	TargetMonkeyIfTrue  int
	TargetMonkeyIfFalse int
	Inspections         int64
	NumberIsOld         bool
}

func getMonkeyBusinessLevel(rounds int, partOne bool) (level int64) {
	monkeys := playMonkeyBusiness(rounds, partOne)
	var inspections []int64
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.Inspections)
	}
	sort.Slice(inspections, func(i, j int) bool { return inspections[i] > inspections[j] })
	level = inspections[0] * inspections[1]
	return level
}

func playMonkeyBusiness(rounds int, divideByThree bool) []*monkey {
	monkeys := getMonkeys()
	var globalModulus int64
	globalModulus = 1

	for _, monkey := range monkeys {
		globalModulus *= monkey.DivisibleBy
	}
	
	for round := 1; round <= rounds; round++ {


		for _, monkey := range monkeys {
			// -- Start with first item
			// 1. Operation -- increase worry level
			// 2. Divide worry level by 3 (only if part one)
			// 3. Test current worry level divisibility
			// 4. Throw to monkey based on divisibility in step 3
			//   a. Remove from beginning of slice for current monkey
			//   b. Add to end of target monkey's item slice

			// If monkey's item list is empty, continue to next monkey
			if len(monkey.StartingItems) < 1 {
				continue
			}

			for _, item := range monkey.StartingItems {
				monkey.Inspections++

				if monkey.NumberIsOld {
					monkey.Operation.Number = item
				}
				operator := monkey.Operation.Operator
				switch operator {
				case "*":
					item *= monkey.Operation.Number
				case "+":
					item += monkey.Operation.Number
				}

				if divideByThree {
					item /= 3
				} else {
					item = item % globalModulus
				}

				var targetMonkey int
				if item%monkey.DivisibleBy == 0 {
					targetMonkey = monkey.TargetMonkeyIfTrue
				} else {
					targetMonkey = monkey.TargetMonkeyIfFalse
				}

				// Append to list of target monkey's items
				monkeys[targetMonkey].StartingItems = append(monkeys[targetMonkey].StartingItems, item)
				// Remove from beginning of current monkey's items
				monkey.StartingItems = monkey.StartingItems[1:]
			}

		}
	}

	return monkeys
}

func getMonkeys() []*monkey {
	input := utils.GetInput("input")

	monkeys := []*monkey{}

	for line := 0; line < len(input)-1; line += 7 {
		startingItemsStrings := strings.Split(input[line+1][18:], ", ")
		var startingItems []int64
		for _, item := range startingItemsStrings {
			itemInt, _ := strconv.ParseInt(item, 10, 64)
			startingItems = append(startingItems, itemInt)
		}

		operator := input[line+2][23:24]
		number, err := strconv.ParseInt(input[line+2][25:], 10, 64)
		var numberIsOld bool
		if err != nil {
			if input[line+2][25:] == "old" {
				numberIsOld = true
			}
		}

		divisibleBy, _ := strconv.ParseInt(input[line+3][21:], 10, 64)

		targetMonkeyIfTrue, _ := strconv.Atoi(strings.Split(input[line+4], "throw to monkey ")[1])
		targetMonkeyIfFalse, _ := strconv.Atoi(strings.Split(input[line+5], "throw to monkey ")[1])

		monkey := monkey{
			StartingItems: startingItems,
			Operation: struct {
				Operator string
				Number   int64
			}{
				Operator: operator,
				Number:   number,
			},
			DivisibleBy:         divisibleBy,
			TargetMonkeyIfTrue:  targetMonkeyIfTrue,
			TargetMonkeyIfFalse: targetMonkeyIfFalse,
			NumberIsOld:         numberIsOld,
		}

		monkeys = append(monkeys, &monkey)
	}

	return monkeys
}
