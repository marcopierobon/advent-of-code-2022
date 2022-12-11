package day6

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithExample(t *testing.T) {
	transmissionSignal := helpers.ReadLines("06_sample_input_part1.txt")
	solution := day6SolutionPart1(transmissionSignal)
	expectedSolution := []int{5, 6, 10, 11}
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	transmissionSignal := []string{}
	solution := day6SolutionPart1(transmissionSignal)
	expectedSolution := []int{}
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithExample(t *testing.T) {
	transmissionSignal := helpers.ReadLines("06_sample_input_part2.txt")
	solution := day6SolutionPart2(transmissionSignal)
	expectedSolution := []int{19, 23, 23, 29, 26}
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithEmpty(t *testing.T) {
	transmissionSignal := []string{}
	solution := day6SolutionPart2(transmissionSignal)
	expectedSolution := []int{}
	assert.Equal(t, expectedSolution, solution)
}
