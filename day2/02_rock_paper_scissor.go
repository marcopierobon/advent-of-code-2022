package day2

import (
	"fmt"
	"helpers"
	"strings"
)

const ROCK_SELECTION_POINTS = 1
const PAPER_SELECTION_POINTS = 2
const SCISSORS_SELECTION_POINTS = 3

const WINNING_POINTS = 6
const DRAWING_POINTS = 3
const LOSING_POINTS = 0

const ROCK_OPPONENT_SELECTION = "A"
const PAPER_OPPONENT_SELECTION = "B"
const SCISSORS_OPPONENT_SELECTION = "C"

const ROCK_MY_SELECTION = "X"
const PAPER_MY_SELECTION = "Y"
const SCISSORS_MY_SELECTION = "Z"

const LOSING_STRATEGY = "X"
const DRAWING_STRATEGY = "Y"
const WINNING_STRATEGY = "Z"

func getSelectionPoints(myCurrentPlay string) int {
	switch myCurrentPlay {
	case ROCK_MY_SELECTION:
		return ROCK_SELECTION_POINTS
	case PAPER_MY_SELECTION:
		return PAPER_SELECTION_POINTS
	case SCISSORS_MY_SELECTION:
		return SCISSORS_SELECTION_POINTS
	default:
		panic(fmt.Sprintf("%s is not an expected value, the allowed ones are X, Y or Z", myCurrentPlay))
	}
}

func getWinningPoints(opponentCurrentPlay, myCurrentPlay string) int {
	switch myCurrentPlay {
	case ROCK_MY_SELECTION:
		switch opponentCurrentPlay {
		case ROCK_OPPONENT_SELECTION:
			return DRAWING_POINTS
		case PAPER_OPPONENT_SELECTION:
			return LOSING_POINTS
		case SCISSORS_OPPONENT_SELECTION:
			return WINNING_POINTS
		default:
			panic(fmt.Sprintf("%s is not an expected value, the allowed ones are A, B or C", opponentCurrentPlay))
		}
	case PAPER_MY_SELECTION:
		switch opponentCurrentPlay {
		case ROCK_OPPONENT_SELECTION:
			return WINNING_POINTS
		case PAPER_OPPONENT_SELECTION:
			return DRAWING_POINTS
		case SCISSORS_OPPONENT_SELECTION:
			return LOSING_POINTS
		default:
			panic(fmt.Sprintf("%s is not an expected value, the allowed ones are A, B or C", opponentCurrentPlay))
		}
	case SCISSORS_MY_SELECTION:
		switch opponentCurrentPlay {
		case ROCK_OPPONENT_SELECTION:
			return LOSING_POINTS
		case PAPER_OPPONENT_SELECTION:
			return WINNING_POINTS
		case SCISSORS_OPPONENT_SELECTION:
			return DRAWING_POINTS
		default:
			panic(fmt.Sprintf("%s is not an expected value, the allowed ones are A, B or C", opponentCurrentPlay))
		}
	default:
		panic(fmt.Sprintf("%s is not an expected value, the allowed ones are X, Y or Z", myCurrentPlay))
	}
}

func determineMyPlay(opponentCurrentPlay, myTargetStrategy string) string {
	switch myTargetStrategy {
	case LOSING_STRATEGY:
		switch opponentCurrentPlay {
		case ROCK_OPPONENT_SELECTION:
			return SCISSORS_MY_SELECTION
		case PAPER_OPPONENT_SELECTION:
			return ROCK_MY_SELECTION
		case SCISSORS_OPPONENT_SELECTION:
			return PAPER_MY_SELECTION
		default:
			panic(fmt.Sprintf("%s is not an expected value, the allowed ones are A, B or C", opponentCurrentPlay))
		}
	//need to draw
	case DRAWING_STRATEGY:
		switch opponentCurrentPlay {
		case ROCK_OPPONENT_SELECTION:
			return ROCK_MY_SELECTION
		case PAPER_OPPONENT_SELECTION:
			return PAPER_MY_SELECTION
		case SCISSORS_OPPONENT_SELECTION:
			return SCISSORS_MY_SELECTION
		default:
			panic(fmt.Sprintf("%s is not an expected value, the allowed ones are A, B or C", opponentCurrentPlay))
		}
	//need to win
	case WINNING_STRATEGY:
		switch opponentCurrentPlay {
		case ROCK_OPPONENT_SELECTION:
			return PAPER_MY_SELECTION
		case PAPER_OPPONENT_SELECTION:
			return SCISSORS_MY_SELECTION
		case SCISSORS_OPPONENT_SELECTION:
			return ROCK_MY_SELECTION
		default:
			panic(fmt.Sprintf("%s is not an expected value, the allowed ones are A, B or C", opponentCurrentPlay))
		}
	default:
		panic(fmt.Sprintf("%s is not an expected value, the allowed ones are X, Y or Z", myTargetStrategy))
	}
}

func calculateScoreStrategy1(strategyPredictionRounds []string) int {
	var totalScore = 0
	for _, currentRoundPlays := range strategyPredictionRounds {
		bothPlayersPlay := strings.Fields(currentRoundPlays)
		opponentCurrentPlay := bothPlayersPlay[0]
		myCurrentPlay := bothPlayersPlay[1]
		totalScore += getSelectionPoints(myCurrentPlay)
		totalScore += getWinningPoints(opponentCurrentPlay, myCurrentPlay)
	}
	return totalScore
}

func calculateScoreStrategy2(strategyPredictionRounds []string) int {
	var totalScore = 0
	for _, currentRoundPlays := range strategyPredictionRounds {
		bothPlayersPlay := strings.Fields(currentRoundPlays)
		opponentCurrentPlay := bothPlayersPlay[0]
		myTargetStrategy := bothPlayersPlay[1]
		myCurrentPlay := determineMyPlay(opponentCurrentPlay, myTargetStrategy)
		totalScore += getSelectionPoints(myCurrentPlay)
		totalScore += getWinningPoints(opponentCurrentPlay, myCurrentPlay)
	}
	return totalScore
}

func Day2Solution() {
	strategyPredictionRounds := helpers.ReadLines("day2/02_input.txt")
	totalScoreStrategy1 := calculateScoreStrategy1(strategyPredictionRounds)

	totalScoreStrategy2 := calculateScoreStrategy2(strategyPredictionRounds)

	println(fmt.Sprintf("The rock paper scissors score with strategy 1 is %d and with strategy 2 is %d",
		totalScoreStrategy1, totalScoreStrategy2))

}
