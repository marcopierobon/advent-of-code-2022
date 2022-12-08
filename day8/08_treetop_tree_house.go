package day8

import (
	"fmt"
	"helpers"
)

const START_OF_THE_PACKET_MARKER_POSITIONS = 4
const START_OF_THE_PACKET_MESSAGE_POSITIONS = 14

func Day8Solution() {
	day8SolutionPart1(helpers.ReadLines("day8/08_input.txt"))
	day8SolutionPart2(helpers.ReadLines("day8/08_input.txt"))
}

func day8SolutionPart1(treeMapLines []string) int {
	visibleTreesCount := countVisibleTrees(treeMapLines)
	fmt.Printf("Number of visible trees: %d\n", visibleTreesCount)
	return visibleTreesCount
}

func day8SolutionPart2(treeMapLines []string) int {
	bestScenicScore := getBestScenicScore(treeMapLines)
	fmt.Printf("Highest score for a single tree: %d\n", bestScenicScore)
	return bestScenicScore
}

func countVisibleTrees(treeMapLines []string) int {
	var visibleTreesCount = 0
	for treeMapLineIndex, treeMapLine := range treeMapLines {
		for treeMapColumnIndex, tree := range treeMapLine {
			if treeMapColumnIndex == 0 || treeMapColumnIndex == len(treeMapLines)-1 ||
				treeMapLineIndex == 0 || treeMapLineIndex == len(treeMapLines)-1 {
				visibleTreesCount++
				continue
			}
			currentTreeHeight := helpers.Rtoi(tree)

			treesToTheLeft := helpers.Atoiv(treeMapLine[0:treeMapColumnIndex])
			treesToTheRight := helpers.Atoiv(treeMapLine[treeMapColumnIndex+1:])
			treesToTheTop := getVerticalSlice(treeMapLines, treeMapColumnIndex)[0:treeMapLineIndex]
			treesToTheBottom := getVerticalSlice(treeMapLines, treeMapColumnIndex)[treeMapLineIndex+1 : len(treeMapLines)]

			isTheTreeVisibleInAnyDirection := isTheTreeVisibleInAnyDirection(currentTreeHeight,
				treesToTheLeft, treesToTheRight, treesToTheTop, treesToTheBottom)

			if isTheTreeVisibleInAnyDirection {
				visibleTreesCount++
			}
		}
	}
	return visibleTreesCount
}

func getBestScenicScore(treeMapLines []string) int {
	var bestScenicScore = 0
	for treeMapLineIndex, treeMapLine := range treeMapLines {
		for treeMapColumnIndex, tree := range treeMapLine {
			if treeMapColumnIndex == 0 || treeMapColumnIndex == len(treeMapLines)-1 ||
				treeMapLineIndex == 0 || treeMapLineIndex == len(treeMapLines)-1 {
				continue
			}
			currentTreeHeight := helpers.Rtoi(tree)

			treesToTheLeft := helpers.Reverse(helpers.Atoiv(treeMapLine[0:treeMapColumnIndex]))
			treesToTheRight := helpers.Atoiv(treeMapLine[treeMapColumnIndex+1:])
			treesToTheTop := helpers.Reverse(getVerticalSlice(treeMapLines, treeMapColumnIndex)[0:treeMapLineIndex])
			treesToTheBottom := getVerticalSlice(treeMapLines, treeMapColumnIndex)[treeMapLineIndex+1 : len(treeMapLines)]

			currentTreeScenicScore := calculateSceniceScore(currentTreeHeight,
				treesToTheLeft, treesToTheRight, treesToTheTop, treesToTheBottom)

			if currentTreeScenicScore > bestScenicScore {
				bestScenicScore = currentTreeScenicScore
			}
		}
	}
	return bestScenicScore
}

func getVerticalSlice(treeMapLines []string, treeMapColumnIndex int) []int {
	var verticalSlice = []int{}
	for _, element := range treeMapLines {
		verticalSlice = append(verticalSlice, helpers.Btoi(element[treeMapColumnIndex]))
	}
	return verticalSlice
}

func calculateSceniceScore(currentTreeHeight int, treesHeightsInAllDirections ...[]int) int {
	var totalScenicScore = 1
	for _, treesHeightsInADirection := range treesHeightsInAllDirections {
		var scenicScoreInCurrentDirection = len(treesHeightsInADirection)
		for treeIndexInCurrentDirection, treeHeightInTheCurrentDirection := range treesHeightsInADirection {
			if treeHeightInTheCurrentDirection >= currentTreeHeight {
				scenicScoreInCurrentDirection = treeIndexInCurrentDirection + 1
				break
			}
		}
		if scenicScoreInCurrentDirection > 0 {
			totalScenicScore *= scenicScoreInCurrentDirection
		}
	}
	return totalScenicScore
}

func isTheTreeVisibleInAnyDirection(currentTreeHeight int, treesHeightsInAllDirections ...[]int) bool {

	for _, treesHeightsInADirection := range treesHeightsInAllDirections {
		var isTheTreeVisibleInTheCurrentDirection = true
		for _, treeHeightInTheCurrentDirection := range treesHeightsInADirection {
			if treeHeightInTheCurrentDirection >= currentTreeHeight {
				isTheTreeVisibleInTheCurrentDirection = false
				break
			}
		}
		if isTheTreeVisibleInTheCurrentDirection {
			return true
		}
	}
	return false
}
