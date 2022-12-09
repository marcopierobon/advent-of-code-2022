package day9

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithExample(t *testing.T) {
	listOfHeadMovements := helpers.ReadLines("09_sample_input.txt")
	solution := day9SolutionPart1(listOfHeadMovements)
	expectedSolution := 13
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithFullData(t *testing.T) {
	listOfHeadMovements := helpers.ReadLines("09_input.txt")
	solution := day9SolutionPart1(listOfHeadMovements)
	expectedSolution := 6339
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithMyInput(t *testing.T) {
	listOfHeadMovements := helpers.ReadLines("09_my_input.txt")
	solution := day9SolutionPart1(listOfHeadMovements)
	expectedSolution := 8
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	treeMap := []string{}
	solution := day9SolutionPart1(treeMap)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithExample(t *testing.T) {
	listOfHeadMovements := helpers.ReadLines("09_sample_input.txt")
	solution := day9SolutionPart2(listOfHeadMovements)
	expectedSolution := 1
	assert.Equal(t, expectedSolution, solution)
}

// func TestPart2WithExample(t *testing.T) {
// 	treeMap := helpers.ReadLines("08_sample_input_part1.txt")
// 	solution := day8SolutionPart2(treeMap)
// 	expectedSolution := 8
// 	assert.Equal(t, expectedSolution, solution)
// }

// func TestPart2WithEmpty(t *testing.T) {
// 	treeMap := []string{}
// 	solution := day8SolutionPart2(treeMap)
// 	expectedSolution := 0
// 	assert.Equal(t, expectedSolution, solution)
// }
