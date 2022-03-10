package main

import "math/rand"

func RockPaperScissor() int {
	return rand.Intn(3)
}

func PlayValue(play int) string {
	playStrings := [3]string{"rock", "paper", "scissor"}
	return playStrings[play]
}

func PlayAgainst(play1 int, play2 int) int {
	if play1 == play2 {
		return 0
	}
	if (play1 == 0 && play2 == 2) || (play1 == 1 && play2 == 0) || (play1 == 2 && play2 == 1) {
		return 1
	}
	if (play1 == 0 && play2 == 1) || (play1 == 1 && play2 == 2) || (play1 == 2 && play2 == 0) {
		return 2
	}
	return -1
}

func WinnerValue(winner int) string {
	winnerStrings := [3]string{"Draw", "You Win", "You Lose"}
	return winnerStrings[winner]
}
