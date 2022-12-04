package day4

import (
	"fmt"
	"helpers"
	"strings"
)

func Day4Solution() {
	day4SolutionPart1(helpers.ReadLines("day4/04_input.txt"))
	day4SolutionPart2(helpers.ReadLines("day4/04_input.txt"))
}

func day4SolutionPart1(backpacksContent []string) int {
	cleaningAssignmentsFullOverlapsCount := getCleaningAssignmentsFullOverlapsCount(backpacksContent)
	println(fmt.Sprintf("The sum of full overlapping spaces in the pair assignments is %d",
		cleaningAssignmentsFullOverlapsCount))
	return cleaningAssignmentsFullOverlapsCount
}

func day4SolutionPart2(backpacksContent []string) int {
	cleaningAssignmentsAnyOverlapsCount := getCleaningAssignmentsAnyOverlapsCount(backpacksContent)
	println(fmt.Sprintf("The sum of any (both full or partial) overlapping spaces in the pair assignments is %d",
		cleaningAssignmentsAnyOverlapsCount))
	return cleaningAssignmentsAnyOverlapsCount
}

func getCleaningAssignmentsAnyOverlapsCount(backpacksContent []string) int {
	var overlapCount = 0
	for index, element := range backpacksContent {
		currentPairAssignments := strings.Split(element, ",")
		isPairAssignmentOverlapping := isPairAssignmentOverlapping(currentPairAssignments)
		if isPairAssignmentOverlapping {
			overlapCount++
		}
		fmt.Printf("At iteration %d the assignments %s and %s are overlapping=%t\n",
			index, currentPairAssignments[0], currentPairAssignments[1], isPairAssignmentOverlapping)
	}
	return overlapCount
}

func getCleaningAssignmentsFullOverlapsCount(backpacksContent []string) int {
	var overlapCount = 0
	for _, element := range backpacksContent {
		currentPairAssignments := strings.Split(element, ",")
		isPairAssignmentOverlapping := isPairAssignmentFullyOverlapping(currentPairAssignments)
		if isPairAssignmentOverlapping {
			overlapCount++
		}
	}
	return overlapCount
}

func isPairAssignmentFullyOverlapping(currentPairAssignments []string) bool {
	firstAssignmentStartAndEndSections := strings.Split(currentPairAssignments[0], "-")
	secondAssignmentStartAndEndSections := strings.Split(currentPairAssignments[1], "-")
	firstAssignmentStartSection := helpers.Atoi(firstAssignmentStartAndEndSections[0])
	firstAssignmentEndSection := helpers.Atoi(firstAssignmentStartAndEndSections[1])
	secondAssignmentStartSection := helpers.Atoi(secondAssignmentStartAndEndSections[0])
	secondAssignmentEndSection := helpers.Atoi(secondAssignmentStartAndEndSections[1])

	return (firstAssignmentStartSection <= secondAssignmentStartSection && firstAssignmentEndSection >= secondAssignmentEndSection) ||
		(secondAssignmentStartSection <= firstAssignmentStartSection && secondAssignmentEndSection >= firstAssignmentStartSection)
}

func isPairAssignmentOverlapping(currentPairAssignments []string) bool {
	firstAssignmentStartAndEndSections := strings.Split(currentPairAssignments[0], "-")
	secondAssignmentStartAndEndSections := strings.Split(currentPairAssignments[1], "-")
	firstAssignmentStartSection := helpers.Atoi(firstAssignmentStartAndEndSections[0])
	firstAssignmentEndSection := helpers.Atoi(firstAssignmentStartAndEndSections[1])
	secondAssignmentStartSection := helpers.Atoi(secondAssignmentStartAndEndSections[0])
	secondAssignmentEndSection := helpers.Atoi(secondAssignmentStartAndEndSections[1])

	return (firstAssignmentStartSection <= secondAssignmentStartSection && firstAssignmentEndSection >= secondAssignmentStartSection) ||
		(secondAssignmentStartSection <= firstAssignmentStartSection && secondAssignmentEndSection >= firstAssignmentStartSection)
}
