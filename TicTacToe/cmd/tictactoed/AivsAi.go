package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/NehemiahG7/project-1/TicTacToe/gameio/ai"
	"github.com/NehemiahG7/project-1/TicTacToe/gamewin"
)

func aivsAi(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("TicTacToe/cmd/tictactoed/web/ttt3.html")
	t.Execute(w, nil)

	gamewin.CheckWin()
	for gamewin.CheckWin() == false {
		hboard(w, r)
		ai.AiXmove()

		gamewin.CheckWin()
		if gamewin.CheckWin() == false {
			hboard(w, r)
			ai.AiOmove()

			gamewin.CheckWin()
		}
	}
	hboard(w, r)
	fmt.Fprintf(w, "<br><p1>%v</p1>", gamewin.WonChat)
	fmt.Println("**************")
	fmt.Println(gamewin.WonChat)
}
