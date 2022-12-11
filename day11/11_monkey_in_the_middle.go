package day11

import (
	"fmt"
	"helpers"
	"strings"
)

const (
	Sum      = 0
	Multiply = 1
	Square   = 2
)

const ROUNDS_QUICK = 20
const ROUNDS_EXTENDED = 10000

const RELIEVE_WORRY_LEVEL = true

func square(x int) int {
	return x * x
}

func mul(y int) func(int) int {
	return func(x int) int {
		return x * y
	}
}

func add(y int) func(int) int {
	return func(x int) int {
		return x + y
	}
}

type Monkey struct {
	worryLevelPerItemForCurrentMonkey []int
	itemsExaminedSoFar                int
	changeWorryLevelFunc              func(int) int
	divisibleByValue                  int
	monkeyNumberToThrowToIfTrue       int
	monkeyNumberToThrowToIfFalse      int
}

func Day11Solution() {
	day11SolutionPart1(helpers.ReadLines("day11/11_input.txt"))
	day11SolutionPart2(helpers.ReadLines("day11/11_input.txt"))
}

func day11SolutionPart1(monkeysChecks []string) int {
	max, secondMax := getMonkeyBusinessLevelOperators(monkeysChecks, RELIEVE_WORRY_LEVEL, ROUNDS_QUICK)
	fmt.Printf("Using a worry relieve strategy, the two largest values are %d and %d. Multiplied together they are: %d\n",
		max, secondMax, max*secondMax)
	return max * secondMax
}

func day11SolutionPart2(monkeysChecks []string) int {
	max, secondMax := getMonkeyBusinessLevelOperators(monkeysChecks, !RELIEVE_WORRY_LEVEL, ROUNDS_EXTENDED)
	fmt.Printf("Without a worry relieve strategy, the two largest values are %d and %d. Multiplied together they are: %d\n",
		max, secondMax, max*secondMax)
	return max * secondMax
}

func getMonkeyBusinessLevelOperators(monkeysChecks []string, isWorryLevelRelieved bool, numberOfRounds int) (int, int) {
	monkeys := getInitialItemsPerMonkey(monkeysChecks)
	activityPerMonkey := [8]int{}
	for i := 0; i < numberOfRounds; i++ {
		mod := 1
		for i := range monkeys {
			mod *= monkeys[i].divisibleByValue
		}
		calculateNextRound(monkeys, isWorryLevelRelieved, mod)
	}

	for j := 0; j < len(monkeys); j++ {
		activityPerMonkey[j] += monkeys[j].itemsExaminedSoFar
	}
	max, secondMax := findTwoLargestValuesOfItemsTouched(activityPerMonkey)
	return max, secondMax
}

func findTwoLargestValuesOfItemsTouched(activityPerMonkey [8]int) (int, int) {
	max := helpers.Max(activityPerMonkey[:])
	var secondMax = 0
	for _, element := range activityPerMonkey {
		if element > secondMax && element != max {
			secondMax = element
		}
	}
	return max, secondMax
}

func calculateNextRound(monkeys []*Monkey, isWorryLevelRelieved bool, mod int) {
	for _, monkey := range monkeys {

		if len(monkey.worryLevelPerItemForCurrentMonkey) == 0 {
			continue
		}
		indexLastInitialElement := len(monkey.worryLevelPerItemForCurrentMonkey)
		for _, worryItem := range monkey.worryLevelPerItemForCurrentMonkey {
			updatedWorry := calculateWorry(worryItem, monkey.changeWorryLevelFunc,
				isWorryLevelRelieved, mod)

			isTestConditionTrue :=
				evaluateTrueTestCondition(updatedWorry, monkey.divisibleByValue)
			if isTestConditionTrue {
				monkeys[monkey.monkeyNumberToThrowToIfTrue].worryLevelPerItemForCurrentMonkey =
					append(monkeys[monkey.monkeyNumberToThrowToIfTrue].worryLevelPerItemForCurrentMonkey, updatedWorry)

			} else {
				monkeys[monkey.monkeyNumberToThrowToIfFalse].worryLevelPerItemForCurrentMonkey =
					append(monkeys[monkey.monkeyNumberToThrowToIfFalse].worryLevelPerItemForCurrentMonkey, updatedWorry)
			}
		}

		monkey.itemsExaminedSoFar += len(monkey.worryLevelPerItemForCurrentMonkey)
		monkey.worryLevelPerItemForCurrentMonkey = monkey.worryLevelPerItemForCurrentMonkey[indexLastInitialElement:]
	}
}

func evaluateTrueTestCondition(currentItemWorry, divisibleByValue int) bool {
	return currentItemWorry%divisibleByValue == 0
}

func calculateWorry(worryItem int, changeWorryLevelFunc func(int) int,
	isWorryLevelRelieved bool, mod int) int {
	var currentItemWorry = 0
	currentItemWorry = changeWorryLevelFunc(worryItem)
	if isWorryLevelRelieved {
		currentItemWorryAfterMonkeyGetsBored := currentItemWorry / 3
		return currentItemWorryAfterMonkeyGetsBored
	} else {
		return currentItemWorry % mod
	}
}

func getInitialItemsPerMonkey(monkeysChecks []string) []*Monkey {
	monkeys := make([]*Monkey, 0)
	for index, line := range monkeysChecks {
		if strings.HasPrefix(line, "Monkey") {
			worryLevelPerItemForCurrentMonkey := getInitialWorryLevelPerItemForCurrentMonkey(monkeysChecks[index+1])
			targetFunc := getWorryOperation(monkeysChecks[index+2])
			divisibleByValue, monkeyNumberToThrowToIfTrue, monkeyNumberToThrowToIfFalse :=
				getThrowingTest(monkeysChecks[index+3 : index+6])
			itemsExaminedSoFar := 0
			var currentMonkey = Monkey{worryLevelPerItemForCurrentMonkey, itemsExaminedSoFar, targetFunc,
				divisibleByValue, monkeyNumberToThrowToIfTrue, monkeyNumberToThrowToIfFalse}
			monkeys = append(monkeys, &currentMonkey)
		}
	}
	return monkeys
}

func getThrowingTest(throwingTestInput []string) (int, int, int) {
	divisibleByValue := helpers.Atoi(strings.Split(throwingTestInput[0], "divisible by ")[1])
	monkeyNumberToThrowToIfTrue := helpers.Atoi(throwingTestInput[1][len(throwingTestInput[1])-1:])
	monkeyNumberToThrowToIfFalse := helpers.Atoi(throwingTestInput[2][len(throwingTestInput[2])-1:])
	return divisibleByValue, monkeyNumberToThrowToIfTrue, monkeyNumberToThrowToIfFalse
}

func getWorryOperation(operationInput string) func(int) int {
	operation := strings.Fields(strings.Split(operationInput, "= old ")[1])

	if operation[1] == "old" {
		return square
	} else {
		factor := helpers.Atoi(operation[1])
		if operation[0] == "+" {
			return add(factor)
		} else if operation[0] == "*" {
			return mul(factor)
		} else {
			panic(fmt.Sprintf("Could not determine the operation for %s\n", operationInput))
		}
	}

}

func getInitialWorryLevelPerItemForCurrentMonkey(startingItemsLine string) []int {
	initialWorryLevelPerItemForCurrentMonkey := make([]int, 0)
	var currentMonkeyItemsString = strings.Split(strings.Split(startingItemsLine, ": ")[1], ", ")
	for _, element := range currentMonkeyItemsString {
		initialWorryLevelPerItemForCurrentMonkey = append(initialWorryLevelPerItemForCurrentMonkey, helpers.Atoi(element))
	}
	return initialWorryLevelPerItemForCurrentMonkey
}
