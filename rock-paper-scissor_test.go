package main

import "testing"

func TestRockPaperScissor(t *testing.T) {
	t.Run("Random value", func(t *testing.T) {
		var gotchas [3]bool
		tries := 9
		for i := 0; i < tries; i++ {
			got := RockPaperScissor()
			gotchas[got] = true
		}

		didProduceAllValidValues := true
		for _, gotcha := range gotchas {
			didProduceAllValidValues = didProduceAllValidValues && gotcha
		}
		if !didProduceAllValidValues {
			t.Errorf("Did not produce 0, 1, and 2 after %v tries", tries)
		}
	})
}

func TestPlayAgainst(t *testing.T) {
	tests := []struct {
		play1  int
		play2  int
		winner int
	}{
		{play1: 0, play2: 1, winner: 2},
		{play1: 1, play2: 2, winner: 2},
		{play1: 2, play2: 0, winner: 2},
		{play1: 0, play2: 0, winner: 0},
		{play1: 1, play2: 1, winner: 0},
		{play1: 2, play2: 2, winner: 0},
		{play1: 0, play2: 2, winner: 1},
		{play1: 1, play2: 0, winner: 1},
		{play1: 2, play2: 1, winner: 1},
	}
	for _, play := range tests {
		playSituation := " (" + PlayValue(play.play1) + " vs " + PlayValue(play.play2) + ")"
		gameExpecation := "Draw"
		if play.winner == 1 {
			gameExpecation = "Player 1 Wins"
		} else if play.winner == 2 {
			gameExpecation = "Player 2 Wins"
		}
		t.Run(gameExpecation+playSituation, func(t *testing.T) {
			gameWinner := PlayAgainst(play.play1, play.play2)
			if gameWinner != play.winner {
				t.Errorf("Incorrect outcome. Expected(%v), Got(%v)", play.winner, gameWinner)
			}
		})
	}
}
