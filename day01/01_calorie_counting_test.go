package day01

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithExample(t *testing.T) {
	backpacksContent := helpers.ReadLines("01_sample_input.txt")
	solution := day1SolutionPart1(backpacksContent)
	expectedSolution := 24000
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	backpacksContent := []string{}
	solution := day1SolutionPart1(backpacksContent)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithExample(t *testing.T) {
	backpacksContent := helpers.ReadLines("01_sample_input.txt")
	solution := day1SolutionPart2(backpacksContent)
	expectedSolution := 45000
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithEmpty(t *testing.T) {
	backpacksContent := []string{}
	solution := day1SolutionPart2(backpacksContent)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}
