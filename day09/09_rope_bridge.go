package day9

import (
	"fmt"
	"helpers"
	"strings"
)

const SHORT_ROPE_LENGTH = 2
const LONG_ROPE_LENGTH = 10

func Day9Solution() {
	day9SolutionPart1(helpers.ReadLines("day09/09_input.txt"))
	day9SolutionPart2(helpers.ReadLines("day09/09_input.txt"))
}

func day9SolutionPart1(listOfHeadMovements []string) int {
	elementsPositions := [][2]int{
		{0, 0},
		{0, 0},
	}
	positionsVisitedByTail := moveRope(listOfHeadMovements, elementsPositions)
	fmt.Printf("The tail visited %d positions using a rope of length %d\n",
		positionsVisitedByTail, len(listOfHeadMovements))
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
	positionsVisitedByTail := moveRope(listOfHeadMovements, elementsPositions)
	fmt.Printf("The tail visited %d positions using a rope of length %d\n",
		positionsVisitedByTail, len(listOfHeadMovements))
	return positionsVisitedByTail
}

func moveRope(listOfHeadMovements []string, elementsPositions [][2]int) int {
	positionCoveredByTail := map[string]bool{}

	positionCoveredByTail[fmt.Sprintf("%d-%d", 0, 0)] = true
	for _, element := range listOfHeadMovements {
		movementsSplit := strings.Fields(element)
		direction := movementsSplit[0]
		numberOfSteps := helpers.Atoi(movementsSplit[1])

		for i := 0; i < numberOfSteps; i++ {
			for j := 0; j < len(elementsPositions)-1; j++ {
				moveRopeNodePairs(direction, &elementsPositions[j], &elementsPositions[j+1], j == 0)
			}
			positionCoveredByTail[fmt.Sprintf("%d-%d",
				elementsPositions[len(elementsPositions)-1][0],
				elementsPositions[len(elementsPositions)-1][1])] = true
		}
	}
	return len(positionCoveredByTail)
}

func moveRopeNodePairs(direction string, leadingElement *[2]int, trailingElement *[2]int, isHead bool) {
	if isHead {
		moveHead(direction, leadingElement)
	}
	moveTheTrailingNode(leadingElement, trailingElement)
}

func moveTheTrailingNode(leadingElement *[2]int, trailingElement *[2]int) {
	horizontalDistance, verticalDistance :=
		getAbsoluteDistances(leadingElement[0], leadingElement[1], trailingElement[0], trailingElement[1])
	if horizontalDistance+verticalDistance == 3 {
		trailingElement[0] += 1 * helpers.Sgn(leadingElement[0]-trailingElement[0])
		trailingElement[1] += 1 * helpers.Sgn(leadingElement[1]-trailingElement[1])
	} else if horizontalDistance == 2 {
		trailingElement[0] += 1 * helpers.Sgn(leadingElement[0]-trailingElement[0])
	} else if verticalDistance == 2 {
		trailingElement[1] += 1 * helpers.Sgn(leadingElement[1]-trailingElement[1])
	}
}

func moveHead(direction string, headNode *[2]int) {
	switch direction {
	case "R":
		{
			headNode[0] = (headNode[0] + 1)
		}
	case "L":
		{
			headNode[0] = (headNode[0] - 1)
		}
	case "U":
		{
			headNode[1] = (headNode[1] + 1)
		}
	case "D":
		{
			headNode[1] = (headNode[1] - 1)
		}
	}
}

func getAbsoluteDistances(firstElementX int, firstElementY int, secondElementX int, secondElementY int) (int, int) {
	var horizontalDistance = helpers.Abs(firstElementX - secondElementX)
	var verticalDistance = helpers.Abs(firstElementY - secondElementY)
	return horizontalDistance, verticalDistance
}
