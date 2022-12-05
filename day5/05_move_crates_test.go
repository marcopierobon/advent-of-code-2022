package day5

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithExample(t *testing.T) {
	initialStacksConfiguration := [][]string{
		{},
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	stacksMovementOperations := helpers.ReadLines("05_sample_input.txt")
	solution := day5SolutionPart1(initialStacksConfiguration, stacksMovementOperations)
	expectedSolution := [][]string{
		{"C"},
		{"M"},
		{"P", "D", "N", "Z"},
	}
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	initialStacksConfiguration := [][]string{}
	stacksMovementOperations := []string{}
	solution := day5SolutionPart1(initialStacksConfiguration, stacksMovementOperations)
	expectedSolution := [][]string{}
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithExample(t *testing.T) {
	initialStacksConfiguration := [][]string{
		{},
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	stacksMovementOperations := helpers.ReadLines("05_sample_input.txt")
	solution := day5SolutionPart2(initialStacksConfiguration, stacksMovementOperations)
	expectedSolution := [][]string{
		{"M"},
		{"C"},
		{"P", "Z", "N", "D"},
	}
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2Empty(t *testing.T) {
	initialStacksConfiguration := [][]string{}
	stacksMovementOperations := []string{}
	solution := day5SolutionPart2(initialStacksConfiguration, stacksMovementOperations)
	expectedSolution := [][]string{}
	assert.Equal(t, expectedSolution, solution)
}
