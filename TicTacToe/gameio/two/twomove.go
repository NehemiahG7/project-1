package two

import (
	"fmt"

	"github.com/NehemiahG7/project-1/TicTacToe/gameboard"
	"github.com/NehemiahG7/project-1/TicTacToe/gamewin"
)

func twomove() {
	goodInput := false
	x := 0
	for goodInput == false {
		fmt.Println("O's move pick 1-9")
		fmt.Scanln(&x)
		for i := 0; i < len(gameboard.Board); i++ {
			if x == (i+1) && gameboard.Board[i].Fill == false {
				gameboard.Board[i].Slogic = "O"
				gameboard.Board[i].Fill = true
				gamewin.Moves++
				goodInput = true
				break
			}
		}

		if goodInput == false {
			fmt.Println("Error: The Number you entered has been chosen or you didn't enter a number.")
			fmt.Print(gameboard.PrintBoard())
		}
	}
	gamewin.CheckWin()
}
