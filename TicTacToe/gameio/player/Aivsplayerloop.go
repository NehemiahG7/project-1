package player

import (
	"fmt"

	"github.com/NehemiahG7/project-1/TicTacToe/gameboard"
	"github.com/NehemiahG7/project-1/TicTacToe/gameio/ai"
	"github.com/NehemiahG7/project-1/TicTacToe/gamewin"
)

//Aiplayer runs the game with ai controlling O.
func Aiplayer() {
Reset:
	again := ""
	gamewin.CheckWin()
	for gamewin.CheckWin() == false {
		fmt.Print(gameboard.PrintBoard())
		Xmove()
		gamewin.CheckWin()
		if gamewin.CheckWin() == false {
			fmt.Print(gameboard.PrintBoard())
			ai.AiOmove()
			gamewin.CheckWin()
		}
	}
	fmt.Print(gameboard.PrintBoard())
	fmt.Println(gamewin.WonChat)
	fmt.Println("Play Again? y/n")

	fmt.Scan(&again)
	if again == "y" || again == "Y" || again == "Yes" || again == "yes" {

		gamewin.ResetWin()
		goto Reset

	}
}
