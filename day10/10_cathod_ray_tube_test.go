package day10

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithBasicExample(t *testing.T) {
	instructions := helpers.ReadLines("10_sample_input_basic.txt")
	solution := day10SolutionPart1(instructions)
	expectedSolution := 2000120
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithAdvancedExample(t *testing.T) {
	instructions := helpers.ReadLines("10_sample_input_advanced.txt")
	solution := day10SolutionPart1(instructions)
	expectedSolution := 13140
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	instructions := []string{}
	solution := day10SolutionPart1(instructions)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithBasicExample(t *testing.T) {
	instructions := helpers.ReadLines("10_sample_input_basic.txt")
	solution := day10SolutionPart2(instructions)
	expectedSolution :=
		`██████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░




`
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithAdvancedExample(t *testing.T) {
	instructions := helpers.ReadLines("10_sample_input_advanced.txt")
	solution := day10SolutionPart2(instructions)
	expectedSolution :=
		`████░░░░████░░░░████░░░░████░░░░████░░░░████░░░░████░░░░████░░░░████░░░░████░░░░
██████░░░░░░██████░░░░░░██████░░░░░░██████░░░░░░██████░░░░░░██████░░░░░░██████░░
████████░░░░░░░░████████░░░░░░░░████████░░░░░░░░████████░░░░░░░░████████░░░░░░░░
██████████░░░░░░░░░░██████████░░░░░░░░░░██████████░░░░░░░░░░██████████░░░░░░░░░░
████████████░░░░░░░░░░░░████████████░░░░░░░░░░░░████████████░░░░░░░░░░░░████████
██████████████░░░░░░░░░░░░░░██████████████░░░░░░░░░░░░░░██████████████░░░░░░░░░░`
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithFullInput(t *testing.T) {
	instructions := helpers.ReadLines("10_input.txt")
	solution := day10SolutionPart2(instructions)
	expectedSolution :=
		`██████░░░░██░░░░██░░░░████░░░░████████░░░░████░░░░░░░░████░░██████░░░░██████░░░░
██░░░░██░░██░░██░░░░██░░░░██░░░░░░░░██░░██░░░░██░░░░░░░░██░░██░░░░██░░██░░░░██░░
██░░░░██░░████░░░░░░██░░░░██░░░░░░██░░░░██░░░░██░░░░░░░░██░░██████░░░░██░░░░██░░
██████░░░░██░░██░░░░████████░░░░██░░░░░░████████░░░░░░░░██░░██░░░░██░░██████░░░░
██░░██░░░░██░░██░░░░██░░░░██░░██░░░░░░░░██░░░░██░░██░░░░██░░██░░░░██░░██░░██░░░░
██░░░░██░░██░░░░██░░██░░░░██░░████████░░██░░░░██░░░░████░░░░██████░░░░██░░░░██░░`
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithEmpty(t *testing.T) {
	instructions := []string{}
	solution := day10SolutionPart2(instructions)
	expectedSolution := `




`
	assert.Equal(t, expectedSolution, solution)
}
