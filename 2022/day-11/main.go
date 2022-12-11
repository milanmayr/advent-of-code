package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

func main() {
	print("After 20 rounds of stuff-slinging simian shenanigans, the level of monkey business is: ")
	println(getMonkeyBusinessLevel())
}

type monkey struct {
	StartingItems []int
	Operation     struct {
		Operator string
		Number   int
	}
	DivisibleBy         int
	TargetMonkeyIfTrue  int
	TargetMonkeyIfFalse int
	Inspections         int
	NumberIsOld         bool
}

func getMonkeyBusinessLevel() (level int) {
	monkeys := playMonkeyBusiness(20)
	var inspections []int
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.Inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	level = inspections[0] * inspections[1]
	return level
}

func playMonkeyBusiness(rounds int) []*monkey {
	monkeys := getMonkeys()

	for round := 1; round <= rounds; round++ {
		for _, monkey := range monkeys {
			// -- Start with first item
			// 1. Operation -- increase worry level
			// 2. Divide worry level by 3
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
				case "/":
					item /= monkey.Operation.Number
				case "+":
					item += monkey.Operation.Number
				case "-":
					item -= monkey.Operation.Number
				}

				item /= 3

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
		var startingItems []int
		for _, item := range startingItemsStrings {
			itemInt, _ := strconv.Atoi(item)
			startingItems = append(startingItems, itemInt)
		}

		operator := input[line+2][23:24]
		number, err := strconv.Atoi(input[line+2][25:])
		var numberIsOld bool
		if err != nil {
			if input[line+2][25:] == "old" {
				numberIsOld = true
			}
		}

		divisibleBy, _ := strconv.Atoi(input[line+3][21:])

		targetMonkeyIfTrue, _ := strconv.Atoi(strings.Split(input[line+4], "throw to monkey ")[1])
		targetMonkeyIfFalse, _ := strconv.Atoi(strings.Split(input[line+5], "throw to monkey ")[1])

		monkey := monkey{
			StartingItems: startingItems,
			Operation: struct {
				Operator string
				Number   int
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
