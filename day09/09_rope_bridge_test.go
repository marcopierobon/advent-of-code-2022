package day9

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithSharedExample(t *testing.T) {
	listOfHeadMovements := helpers.ReadLines("09_sample_input_common.txt")
	solution := day9SolutionPart1(listOfHeadMovements)
	expectedSolution := 13
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithCustomInput(t *testing.T) {
	listOfHeadMovements := helpers.ReadLines("09_custom_input.txt")
	solution := day9SolutionPart1(listOfHeadMovements)
	expectedSolution := 8
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	treeMap := []string{}
	solution := day9SolutionPart1(treeMap)
	expectedSolution := 1
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithSharedExample(t *testing.T) {
	listOfHeadMovements := helpers.ReadLines("09_sample_input_common.txt")
	solution := day9SolutionPart2(listOfHeadMovements)
	expectedSolution := 1
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithPart2Example(t *testing.T) {
	listOfHeadMovements := helpers.ReadLines("09_sample_input_part2.txt")
	solution := day9SolutionPart2(listOfHeadMovements)
	expectedSolution := 36
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithEmpty(t *testing.T) {
	treeMap := []string{}
	solution := day9SolutionPart2(treeMap)
	expectedSolution := 1
	assert.Equal(t, expectedSolution, solution)
}
