package main

import (
	"fmt"

	"github.com/vaibhawj/tictactoe/game"
	"github.com/vaibhawj/tictactoe/player"
)

func main() {
	fmt.Println("Welcome to tic tac toe")

	player1 := game.InitializePlayer(1, "X")
	player2 := game.InitializePlayer(2, "O")
	players := [2]player.Player{player1, player2}

	gameSession := game.NewGame(players)
	winner := gameSession.Start()

	fmt.Println("Game over!")
	if winner == "" {
		fmt.Println("Its a tie")
	} else {
		fmt.Printf("%v is the winner\n", winner)
	}
}
