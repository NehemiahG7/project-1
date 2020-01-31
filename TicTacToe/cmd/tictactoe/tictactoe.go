package main

import (
	"github.com/NehemiahG7/project-1/TicTacToe/cmd/tictactoe/config"
	"github.com/NehemiahG7/project-1/TicTacToe/gameio/ai"
	"github.com/NehemiahG7/project-1/TicTacToe/gameio/player"
	"github.com/NehemiahG7/project-1/TicTacToe/gameio/two"
)

func main() {
	switch {
	case config.Tp == true:
		two.Twoplayer()
	case config.Av == true:
		ai.AivsAi()
	default:
		player.Aiplayer()
	}

}
