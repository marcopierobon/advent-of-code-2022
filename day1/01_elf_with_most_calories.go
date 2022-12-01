package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
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

func findMax(caloriesByElf []int) int {
	var max = 0
	for _, value := range caloriesByElf {
		if max < value {
			max = value
		}
	}
	return max
}

func orderDescending(caloriesByElf []int) []int {
	sort.Slice(caloriesByElf, func(i, j int) bool {
		return caloriesByElf[i] > caloriesByElf[j]
	})
	return caloriesByElf
}

func main() {
	itemCalories := readLines("01_input.txt")
	caloriesByElf := groupCaloriesByElf(itemCalories)

	elvesOrderedByCaloriesDesc := orderDescending(caloriesByElf)
	println(fmt.Sprintf("The 3 elves with the most calories have %d, %d and %d respectively",
		elvesOrderedByCaloriesDesc[0], elvesOrderedByCaloriesDesc[1], elvesOrderedByCaloriesDesc[2]))
}
