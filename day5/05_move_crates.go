package day5

import (
	"fmt"
	"helpers"
	"strings"
)

func Day5Solution() {

	var initialStacksConfiguration = initStacks()
	day5SolutionPart1(initialStacksConfiguration, helpers.ReadLines("day5/05_input.txt"))

	initialStacksConfiguration = initStacks()
	day5SolutionPart2(initialStacksConfiguration, helpers.ReadLines("day5/05_input.txt"))
}

func initStacks() [][]string {
	return [][]string{
		{},
		{"W", "D", "G", "B", "H", "R", "V"},
		{"J", "N", "G", "C", "R", "F"},
		{"L", "S", "F", "H", "D", "N", "J"},
		{"J", "D", "S", "V"},
		{"S", "H", "D", "R", "Q", "W", "N", "V"},
		{"P", "G", "H", "C", "M"},
		{"F", "J", "B", "G", "L", "Z", "H", "C"},
		{"S", "J", "R"},
		{"L", "G", "S", "R", "B", "N", "V", "M"},
	}
}

func day5SolutionPart1(initialStacksConfiguration [][]string, cratesMovementOperations []string) [][]string {
	finalStacksConfiguration := performPlannedIndividualCratesMovements(initialStacksConfiguration, cratesMovementOperations)
	if len(finalStacksConfiguration) > 0 {
		finalStacksConfiguration = finalStacksConfiguration[1:]
	}

	var finalConfigurationTopElements = ""
	for i := 0; i < len(finalStacksConfiguration); i++ {
		finalConfigurationTopElements += finalStacksConfiguration[i][len(finalStacksConfiguration[i])-1]
	}

	fmt.Printf("Elements on top of each stack moving crates individually: %s\n", finalConfigurationTopElements)
	return finalStacksConfiguration
}

func day5SolutionPart2(initialStacksConfiguration [][]string, cratesMovementOperations []string) [][]string {
	finalStacksConfiguration := performMultiplePlannedCratesMovements(initialStacksConfiguration, cratesMovementOperations)
	if len(finalStacksConfiguration) > 0 {
		finalStacksConfiguration = finalStacksConfiguration[1:]
	}

	var finalConfigurationTopElements = ""
	for i := 0; i < len(finalStacksConfiguration); i++ {
		if len(finalStacksConfiguration[i]) == 0 {
			finalConfigurationTopElements += "__EMPTY__"
		} else {
			finalConfigurationTopElements += finalStacksConfiguration[i][len(finalStacksConfiguration[i])-1]
		}
	}

	fmt.Printf("Elements on top of each stack  moving crates together: %s\n", finalConfigurationTopElements)
	return finalStacksConfiguration
}

func performPlannedIndividualCratesMovements(initialStacksConfiguration [][]string, cratesMovementOperations []string) [][]string {
	var finalStacksConfiguration = initialStacksConfiguration
	for _, element := range cratesMovementOperations {
		finalStacksConfiguration = performInvidualCrateMovement(element, finalStacksConfiguration)
	}
	return finalStacksConfiguration
}

func performMultiplePlannedCratesMovements(initialStacksConfiguration [][]string, cratesMovementOperations []string) [][]string {
	var finalStacksConfiguration = initialStacksConfiguration
	for _, element := range cratesMovementOperations {
		finalStacksConfiguration = performMultipleCratesMovement(element, finalStacksConfiguration)
	}
	return finalStacksConfiguration
}

func performMultipleCratesMovement(element string, finalStacksConfiguration [][]string) [][]string {
	stringsSeparatedByFrom := strings.Split(element, " from ")
	numberOfCratesToMove := helpers.Atoi(strings.Replace(stringsSeparatedByFrom[0], "move ", "", 1))
	stringsSeparatedByTo := strings.Split(stringsSeparatedByFrom[1], " to ")
	initialStackToMoveFrom := helpers.Atoi(stringsSeparatedByTo[0])
	targetStackToMoveTo := helpers.Atoi(stringsSeparatedByTo[1])
	var updatedStacksConfiguration = finalStacksConfiguration
	currentInitialStackLength := len(updatedStacksConfiguration[initialStackToMoveFrom])
	updatedStacksConfiguration[targetStackToMoveTo] =
		appendAll(updatedStacksConfiguration[targetStackToMoveTo], updatedStacksConfiguration[initialStackToMoveFrom][currentInitialStackLength-numberOfCratesToMove:currentInitialStackLength])

	updatedStacksConfiguration[initialStackToMoveFrom] = updatedStacksConfiguration[initialStackToMoveFrom][:len(updatedStacksConfiguration[initialStackToMoveFrom])-numberOfCratesToMove]

	return updatedStacksConfiguration
}

func appendAll(s1, s2 []string) []string {
	return append(s1, s2...)
}

func performInvidualCrateMovement(element string, finalStacksConfiguration [][]string) [][]string {
	stringsSeparatedByFrom := strings.Split(element, " from ")
	numberOfCratesToMove := helpers.Atoi(strings.Replace(stringsSeparatedByFrom[0], "move ", "", 1))
	stringsSeparatedByTo := strings.Split(stringsSeparatedByFrom[1], " to ")
	initialStackToMoveFrom := helpers.Atoi(stringsSeparatedByTo[0])
	targetStackToMoveTo := helpers.Atoi(stringsSeparatedByTo[1])
	var updatedStacksConfiguration = finalStacksConfiguration
	for i := 0; i < numberOfCratesToMove; i++ {

		currentInitialStackLength := len(updatedStacksConfiguration[initialStackToMoveFrom])

		updatedStacksConfiguration[targetStackToMoveTo] = append(updatedStacksConfiguration[targetStackToMoveTo],
			updatedStacksConfiguration[initialStackToMoveFrom][currentInitialStackLength-1])

		updatedStacksConfiguration[initialStackToMoveFrom] = updatedStacksConfiguration[initialStackToMoveFrom][:len(updatedStacksConfiguration[initialStackToMoveFrom])-1]
	}

	return updatedStacksConfiguration
}
