package day10

import (
	"helpers"
	"strings"
)

const CRT_WIDTH = 40
const CRT_HEIGHT = 6

func Day10Solution() {
	day10SolutionPart1(helpers.ReadLines("day10/10_input.txt"))
	day10SolutionPart2(helpers.ReadLines("day10/10_input.txt"))
}

func day10SolutionPart1(instructions []string) int {
	signalStrength, _ := findSignalStrength(instructions)

	return signalStrength
}

func day10SolutionPart2(instructions []string) string {
	_, charactersPrint := findSignalStrength(instructions)

	return charactersPrint
}

func findSignalStrength(instructions []string) (int, string) {
	var cycleCounter = 0
	var signalStrength = 0
	var registryXValue = 1
	var cycleMod = 20

	resultingPrint := [CRT_WIDTH][CRT_HEIGHT]string{}
	for _, instruction := range instructions {
		instructionFields := strings.Fields(instruction)
		command := instructionFields[0]
		switch command {
		case "addx":
			{
				valueToAdd := helpers.Atoi(instructionFields[1])
				for instructionLength := 0; instructionLength < 2; instructionLength++ {
					cycleCounter++
					updateScreen(cycleCounter, registryXValue, &resultingPrint)
					updateSignalStrength(&cycleMod, &signalStrength, cycleCounter, registryXValue)
				}
				registryXValue += valueToAdd
			}
		case "noop":
			{
				cycleCounter++
				updateScreen(cycleCounter, registryXValue, &resultingPrint)
				updateSignalStrength(&cycleMod, &signalStrength, cycleCounter, registryXValue)
			}

		}
	}
	resultingString := ""

	for i := 0; i < CRT_HEIGHT; i++ {
		for j := 0; j < CRT_WIDTH; j++ {
			resultingString += resultingPrint[j][i]
		}
		if i < CRT_HEIGHT-1 {
			resultingString += "\n"
		}
	}
	return signalStrength, resultingString
}

func updateSignalStrength(cycleMod *int, signalStrength *int, cycleCounter int, registryXValue int) int {
	if cycleCounter%(*cycleMod) == 0 && cycleCounter > 0 {
		*cycleMod += 40
		*signalStrength += cycleCounter * registryXValue
	}
	return *cycleMod
}

func updateScreen(cycleCounter int, registryXValue int, resultingPrint *[CRT_WIDTH][CRT_HEIGHT]string) {
	cycleCounter--
	if cycleCounter > 239 {
		return
	}
	if helpers.Abs(registryXValue-(cycleCounter%40)) <= 1 {
		resultingPrint[cycleCounter%40][(cycleCounter / 40)] = "██"
	} else {
		resultingPrint[cycleCounter%40][(cycleCounter / 40)] = "░░"
	}
}
