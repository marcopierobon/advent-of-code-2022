package day11

import (
	"helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1WithBasicExample(t *testing.T) {
	monkeysChecks := helpers.ReadLines("11_sample_input_basic.txt")
	solution := day11SolutionPart1(monkeysChecks)
	expectedSolution := 10605
	assert.Equal(t, expectedSolution, solution)
}

func TestPart1WithEmpty(t *testing.T) {
	monkeysChecks := []string{}
	solution := day11SolutionPart1(monkeysChecks)
	expectedSolution := 0
	assert.Equal(t, expectedSolution, solution)
}

func TestPart2WithBasicExample(t *testing.T) {
	monkeysChecks := helpers.ReadLines("11_sample_input_basic.txt")
	solution := day11SolutionPart2(monkeysChecks)
	expectedSolution := 2713310158
	assert.Equal(t, expectedSolution, solution)
}
