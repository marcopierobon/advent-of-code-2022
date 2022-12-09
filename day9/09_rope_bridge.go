package day9

import (
	"fmt"
	"helpers"
	"strings"
)

const SHORT_ROPE_LENGTH = 2

const LONG_ROPE_LENGTH = 10

func Day9Solution() {
	day9SolutionPart1(helpers.ReadLines("day9/09_input.txt"))
	// day8SolutionPart2(helpers.ReadLines("day8/08_input.txt"))
}

func day9SolutionPart1(listOfHeadMovements []string) int {
	elementsPositions := [][2]int{
		{0, 0},
		{0, 0},
	}
	positionsVisitedByTail := movePositions(listOfHeadMovements, elementsPositions)
	fmt.Printf("The tail visited %d positions", positionsVisitedByTail)
	return positionsVisitedByTail
}

func day9SolutionPart2(listOfHeadMovements []string) int {
	elementsPositions := [][2]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}
	positionsVisitedByTail := movePositions(listOfHeadMovements, elementsPositions)
	fmt.Printf("The tail visited %d positions", positionsVisitedByTail)
	return positionsVisitedByTail
}

func movePositions(listOfHeadMovements []string, elementsPositions [][2]int) int {
	// var elementsPositions[0][0] = 0
	// var elementsPositions[0][1] = 0
	// var elementsPositions[1][0] = 0
	// var elementsPositions[1][1] = 0
	var totalMovements = 0
	var minX = 1000000
	var maxX = -1000000
	var minY = 1000000
	var maxY = -1000000
	// elementsPositions := [][]int{
	// 	{0, 0},
	// 	{0, 0},
	// 	// {0, 0},
	// 	// {0, 0},
	// 	// {0, 0},
	// 	// {0, 0},
	// 	// {0, 0},
	// 	// {0, 0},
	// 	// {0, 0},
	// 	// {0, 0},
	// }
	positionCoveredByTail := map[string]bool{}

	positionCoveredByTail[fmt.Sprintf("%d-%d", 0, 0)] = true
	for _, element := range listOfHeadMovements {
		movementsSplit := strings.Fields(element)
		direction := movementsSplit[0]
		numberOfSteps := helpers.Atoi(movementsSplit[1])

		fmt.Printf("\n\n\nMoving %s\n", element)

		for i := 0; i < numberOfSteps; i++ {
			totalMovements++
			fmt.Printf("Movement %d/%d\n", i+1, numberOfSteps)

			// fmt.Print("\n")
			// fmt.Printf("Previous value HEAD (%d, %d)", elementsPositions[0][0], elementsPositions[0][1])
			// fmt.Printf("Previous value TAIL (%d, %d)", elementsPositions[1][0], elementsPositions[1][1])
			// fmt.Print("\n")
			for j := 0; j < len(elementsPositions)-1; j++ {
				move(direction, &elementsPositions[j], &elementsPositions[j+1], j == 0)

			}

			positionCoveredByTail[fmt.Sprintf("%d-%d", elementsPositions[len(elementsPositions)-1][0], elementsPositions[len(elementsPositions)-1][1])] = true

			// if elementsPositions[1][0] < minX {
			// 	minX = elementsPositions[1][0]
			// }
			// if elementsPositions[1][0] > maxX {
			// 	maxX = elementsPositions[1][0]
			// }
			// if elementsPositions[1][1] < minY {
			// 	minY = elementsPositions[1][1]
			// }
			// if elementsPositions[1][1] > maxY {
			// 	maxY = elementsPositions[1][1]
			// }

			fmt.Print("\n")
			fmt.Printf("Next value HEAD (%d, %d)", elementsPositions[0][0], elementsPositions[0][1])
			fmt.Printf("Next value TAIL (%d, %d)", elementsPositions[1][0], elementsPositions[1][1])
			fmt.Print("\n")
			// for height := 70; height >= -150; height-- {
			// 	var isToBePrinted = false
			// 	for width := -172; width < 98; width++ {
			// 		if elementsPositions[0][0] > (width-5) && elementsPositions[0][0] < (width+5) && (elementsPositions[0][1] > (height-5) && elementsPositions[0][0] < (height+5)) {
			// 			isToBePrinted = true
			// 			if elementsPositions[0][0] == width && elementsPositions[0][1] == height {
			// 				fmt.Print("H")
			// 			} else if elementsPositions[1][0] == width && elementsPositions[1][1] == height {
			// 				fmt.Print("T")
			// 				// positionCoveredByTail[elementsPositions[1][0]][elementsPositions[1][1]] = true
			// 			} else if width == 0 && height == 0 {
			// 				fmt.Print("s")
			// 			} else {
			// 				fmt.Print(".")
			// 			}
			// 		}
			// 	}
			// 	if isToBePrinted {
			// 		fmt.Print("\n")
			// 	}
			// }
			// fmt.Printf("PositionCoveredByTail %d %v\n", len(positionCoveredByTail), positionCoveredByTail)
		}

	}

	// var minX = -1000000
	// var maxX = 1000000
	// var minY = -1000000
	// var maxY = 1000000

	fmt.Printf("total movements: %d\n", totalMovements)
	fmt.Printf("minX: %d; maxX: %d\n", minX, maxX)
	fmt.Printf("MinY: %d; MaxY: %d\n", minY, maxY)
	return len(positionCoveredByTail)
}

func move(direction string, leadingElement *[2]int, trailingElement *[2]int, isHead bool) {
	switch direction {
	case "R":
		{
			if isHead {
				leadingElement[0] = (leadingElement[0] + 1)
			}
			absoluteDistance := getAbsoluteDistance(leadingElement[0], leadingElement[1], trailingElement[0], trailingElement[1])
			if absoluteDistance == 2 && leadingElement[1] == trailingElement[1] {
				trailingElement[0] = (trailingElement[0] + 1)
			} else if absoluteDistance >= 3 {
				trailingElement[1] = leadingElement[1]
				trailingElement[0] = (trailingElement[0] + 1)
			}
		}
	case "L":
		{
			if isHead {
				leadingElement[0] = (leadingElement[0] - 1)
			}
			absoluteDistance := getAbsoluteDistance(leadingElement[0], leadingElement[1], trailingElement[0], trailingElement[1])
			if absoluteDistance == 2 && leadingElement[1] == trailingElement[1] {
				trailingElement[0] = (trailingElement[0] - 1)
			} else if absoluteDistance >= 3 {
				trailingElement[1] = leadingElement[1]
				trailingElement[0] = (trailingElement[0] - 1)
			}
		}
	case "U":
		{
			if isHead {
				leadingElement[1] = (leadingElement[1] + 1)
			}
			absoluteDistance := getAbsoluteDistance(leadingElement[0], leadingElement[1], trailingElement[0], trailingElement[1])
			if absoluteDistance == 2 && leadingElement[0] == trailingElement[0] {
				trailingElement[1] = (trailingElement[1] + 1)
			} else if absoluteDistance == 3 {
				trailingElement[0] = leadingElement[0]
				trailingElement[1] = (trailingElement[1] + 1)
			}
		}
	case "D":
		{
			if isHead {
				leadingElement[1] = (leadingElement[1] - 1)
			}
			absoluteDistance := getAbsoluteDistance(leadingElement[0], leadingElement[1], trailingElement[0], trailingElement[1])
			if absoluteDistance == 2 && leadingElement[0] == trailingElement[0] {
				trailingElement[1] = (trailingElement[1] - 1)
			} else if absoluteDistance == 3 {
				trailingElement[0] = leadingElement[0]
				trailingElement[1] = (trailingElement[1] - 1)
			}
		}
	}
}

func getAbsoluteDistance(firstElementX int, firstElementY int, secondElementX int, secondElementY int) int {
	var horizontalDistance = helpers.Abs(firstElementX - secondElementX)
	var verticalDistance = helpers.Abs(firstElementY - secondElementY)

	// if horizontalDistance == GRID_HORIZONTAL_SIZE-1 {
	// 	horizontalDistance = 1
	// }

	// if verticalDistance == GRID_VERTICAL_SIZE-1 {
	// 	verticalDistance = 1
	// }

	fmt.Printf("Current distance: %d\n", horizontalDistance+verticalDistance)
	return (horizontalDistance + verticalDistance)
}
