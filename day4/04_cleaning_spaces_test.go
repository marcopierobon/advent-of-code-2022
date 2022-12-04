package day4

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithExample(t *testing.T) {
	cleaningSpacesAssignments := helpers.ReadLines("04_sample_input.txt")
	solution := day4SolutionPart1(cleaningSpacesAssignments)
	expectedSolution := 2
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	cleaningSpacesAssignments := []string{}
	solution := day4SolutionPart1(cleaningSpacesAssignments)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithExample(t *testing.T) {
	cleaningSpacesAssignments := helpers.ReadLines("04_sample_input.txt")
	solution := day4SolutionPart2(cleaningSpacesAssignments)
	expectedSolution := 4
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithEmpty(t *testing.T) {
	cleaningSpacesAssignments := []string{}
	solution := day4SolutionPart2(cleaningSpacesAssignments)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}
