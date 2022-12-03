package day3

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithExample(t *testing.T) {

	backpacksContent := helpers.ReadLines("03_sample_input_part1.txt")
	solution := day3SolutionPart1(backpacksContent)
	expectedSolution := 157
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	backpacksContent := []string{}
	solution := day3SolutionPart1(backpacksContent)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithExample(t *testing.T) {

	backpacksContent := helpers.ReadLines("03_sample_input_part2.txt")
	solution := day3SolutionPart2(backpacksContent)
	expectedSolution := 70
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithEmpty(t *testing.T) {
	backpacksContent := []string{}
	solution := day3SolutionPart2(backpacksContent)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}
