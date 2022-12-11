package day8

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithExample(t *testing.T) {
	treeMap := helpers.ReadLines("08_sample_input_part1.txt")
	solution := day8SolutionPart1(treeMap)
	expectedSolution := 21
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	treeMap := []string{}
	solution := day8SolutionPart1(treeMap)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithExample(t *testing.T) {
	treeMap := helpers.ReadLines("08_sample_input_part1.txt")
	solution := day8SolutionPart2(treeMap)
	expectedSolution := 8
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithEmpty(t *testing.T) {
	treeMap := []string{}
	solution := day8SolutionPart2(treeMap)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}
