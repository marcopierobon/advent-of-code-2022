package day3

import (
	"fmt"
	"helpers"
	"strings"
)

const LOWER_CASE_A_IN_ASCII = 97
const LOWER_CASE_Z_IN_ASCII = 122
const UPPER_CASE_A_IN_ASCII = 65
const UPPER_CASE_Z_IN_ASCII = 90
const PRIORITY_OFFSET_IN_LOWER_CASE = 96
const PRIORITY_OFFSET_IN_UPPER_CASE = 38

func findCommonElement(firstBackPackElements, secondBackPackElements string) byte {
	elementsInFirstBackpacks := make(map[string]bool)
	for _, charValue := range firstBackPackElements {
		elementsInFirstBackpacks[string(charValue)] = true
	}
	for index := range firstBackPackElements {
		if _, ok := elementsInFirstBackpacks[string(secondBackPackElements[index])]; ok {
			matchingElem := secondBackPackElements[index]
			return matchingElem
		}
	}

	panicMessage := fmt.Sprintf("Could not find a common item between %s and %s", firstBackPackElements, secondBackPackElements)
	panic(panicMessage)
}

func findCommonBadgeElement(firstBackPack string, secondBackPack string, thirdBackPack string) byte {
	commonElements := make(map[string]bool)
	for _, charValue := range firstBackPack {
		commonElements[string(charValue)] = true
	}
	for charValue := range commonElements {
		if !(strings.Contains(secondBackPack, charValue)) ||
			!(strings.Contains(thirdBackPack, charValue)) {
			delete(commonElements, string(charValue))
		} else {
			return secondBackPack[strings.Index(secondBackPack, charValue)]
		}
	}
	panicMessage := fmt.Sprintf("Could not find a single item among %s, %s  and %s",
		firstBackPack, secondBackPack, thirdBackPack)
	panic(panicMessage)
}

func calculatePriority(commonElementAscii int) int {
	if commonElementAscii >= LOWER_CASE_A_IN_ASCII && commonElementAscii <= LOWER_CASE_Z_IN_ASCII {
		return commonElementAscii - PRIORITY_OFFSET_IN_LOWER_CASE
	}
	if commonElementAscii >= UPPER_CASE_A_IN_ASCII && commonElementAscii <= UPPER_CASE_Z_IN_ASCII {
		return commonElementAscii - PRIORITY_OFFSET_IN_UPPER_CASE
	}
	panicMessage := fmt.Sprintf("Could not calculate priority for the letter %s, ASCII code %d",
		fmt.Sprint(commonElementAscii), commonElementAscii)
	panic(panicMessage)
}

func calculatePriorityOfDuplicatedElements(backpacksContent []string) int {
	var prioritiesSum = 0
	for _, element := range backpacksContent {
		firstBackPackElements := element[0 : len(element)/2]
		secondBackPackElements := element[len(element)/2:]
		commonElement := findCommonElement(firstBackPackElements, secondBackPackElements)
		currentCommonElementPriority := calculatePriority(int(commonElement))
		prioritiesSum += currentCommonElementPriority

	}
	return prioritiesSum
}

func calculatePriorityOfGroupBadgesElements(backpacksContent []string) int {
	var prioritiesSum = 0
	for i := 0; i < len(backpacksContent)-1; i += 3 {
		commonElement := findCommonBadgeElement(backpacksContent[i], backpacksContent[i+1], backpacksContent[i+2])
		currentCommonElementPriority := calculatePriority(int(commonElement))
		prioritiesSum += currentCommonElementPriority

	}
	return prioritiesSum
}

func Day3Solution() {
	day3SolutionPart1(helpers.ReadLines("day3/03_input.txt"))
	day3SolutionPart2(helpers.ReadLines("day3/03_input.txt"))
}

func day3SolutionPart1(backpacksContent []string) int {
	priorityOfDuplicatedElements := calculatePriorityOfDuplicatedElements(backpacksContent)
	println(fmt.Sprintf("The sum of the priorities of common elements in the backpacks is %d",
		priorityOfDuplicatedElements))
	return priorityOfDuplicatedElements
}

func day3SolutionPart2(backpacksContent []string) int {
	priorityOfDuplicatedElements := calculatePriorityOfGroupBadgesElements(backpacksContent)
	println(fmt.Sprintf("The sum of the priorities of for the group badges is %d",
		priorityOfDuplicatedElements))
	return priorityOfDuplicatedElements
}
