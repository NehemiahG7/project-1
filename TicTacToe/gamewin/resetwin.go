package gamewin

import "github.com/NehemiahG7/project-1/TicTacToe/gameboard"

// ResetWin resets the win states.
func ResetWin() {
	for i := range Win {
		Moves = 0
		WonChat = ""
		Win[i] = false
		gameboard.ResetBoard()

	}
}
