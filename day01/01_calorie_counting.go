package day01

import (
	"fmt"
	"helpers"
	"sort"
	"strconv"
)

func Day1Solution() {
	itemCalories := helpers.ReadLines("day01/01_input.txt")
	day1SolutionPart1(itemCalories)
	day1SolutionPart2(itemCalories)
}

func day1SolutionPart1(itemCalories []string) int {
	caloriesByElf := groupCaloriesByElf(itemCalories)

	elvesOrderedByCaloriesDesc := orderDescending(caloriesByElf)
	fmt.Printf("The elf with the most calories has %d\n",
		elvesOrderedByCaloriesDesc[0])
	return elvesOrderedByCaloriesDesc[0]
}

func day1SolutionPart2(itemCalories []string) int {
	caloriesByElf := groupCaloriesByElf(itemCalories)
	if len(caloriesByElf) < 3 {
		return 0
	}
	elvesOrderedByCaloriesDesc := orderDescending(caloriesByElf)
	topThreeCarriersSum := elvesOrderedByCaloriesDesc[0] + elvesOrderedByCaloriesDesc[1] + elvesOrderedByCaloriesDesc[2]
	fmt.Printf("The 3 elves with the most calories have are carrying in total %d\n", topThreeCarriersSum)
	return topThreeCarriersSum
}

func groupCaloriesByElf(caloriesList []string) []int {
	caloriesPerElf := make([]int, 0)
	var currentElf = 0
	caloriesPerElf = append(caloriesPerElf, 0)

	for _, currentItemCaloriesString := range caloriesList {
		currentItemCaloriesInt, err := strconv.Atoi(currentItemCaloriesString)
		if err != nil {
			currentElf++
			caloriesPerElf = append(caloriesPerElf, 0)
		} else {
			caloriesPerElf[currentElf] += currentItemCaloriesInt
		}
	}
	return caloriesPerElf
}

func orderDescending(caloriesByElf []int) []int {
	sort.Slice(caloriesByElf, func(i, j int) bool {
		return caloriesByElf[i] > caloriesByElf[j]
	})
	return caloriesByElf
}
