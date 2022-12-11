package day6

import (
	"fmt"
	"helpers"
)

const START_OF_THE_PACKET_MARKER_POSITIONS = 4
const START_OF_THE_PACKET_MESSAGE_POSITIONS = 14

func Day6Solution() {
	day6SolutionPart1(helpers.ReadLines("day06/06_input.txt"))
	day6SolutionPart2(helpers.ReadLines("day06/06_input.txt"))
}

func day6SolutionPart1(transmissionSignal []string) []int {
	startOfThePacketMarkerPositions := getStartOfThePacketMarkerPositions(transmissionSignal,
		START_OF_THE_PACKET_MARKER_POSITIONS)
	fmt.Printf("Start of the packet marker position: %v\n", startOfThePacketMarkerPositions)
	return startOfThePacketMarkerPositions
}

func day6SolutionPart2(transmissionSignal []string) []int {
	startOfThePacketMessagePositions := getStartOfThePacketMarkerPositions(transmissionSignal,
		START_OF_THE_PACKET_MESSAGE_POSITIONS)
	fmt.Printf("Start of the packet marker position: %v\n", startOfThePacketMessagePositions)
	return startOfThePacketMessagePositions
}

func getStartOfThePacketMarkerPositions(transmissionSignals []string, numbeOfPositionsToCheck int) []int {
	var startOfThePacketMarketPositions = make([]int, 0)
	for _, transmissionSignal := range transmissionSignals {
		for index := range transmissionSignal {
			if len(transmissionSignal) < index+numbeOfPositionsToCheck-1 {
				panicMessage := fmt.Sprintf("Could not find a solution, stopping at index: %d\n", index)
				panic(panicMessage)
			}
			if areAllElementsDifferent(transmissionSignal[index : index+numbeOfPositionsToCheck]) {
				startOfThePacketMarketPositions = append(startOfThePacketMarketPositions, index+numbeOfPositionsToCheck)
				break
			}
		}
	}
	return startOfThePacketMarketPositions
}

func areAllElementsDifferent(elements string) bool {
	mapWithExistingElements := make(map[string]bool)
	for _, element := range elements {
		if _, ok := mapWithExistingElements[string(element)]; ok {
			return false
		} else {
			mapWithExistingElements[string(element)] = true
		}
	}
	return true
}
