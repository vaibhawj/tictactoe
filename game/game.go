package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vaibhawj/tictactoe/grid"
	"github.com/vaibhawj/tictactoe/player"
)

func InitializePlayer(i int, symbol string) player.Player {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter name of player %v\n", i)
	input, _ := reader.ReadString('\n')
	name := strings.Replace(input, "\n", "", 1)
	return player.NewPlayer(symbol, name)
}

type game struct {
	grid    grid.Grid
	players [2]player.Player
}

func NewGame(players [2]player.Player) *game {
	return &game{
		*grid.NewGrid(),
		players}
}

func (game game) winnerFromSymbol(symbol string) string {
	players := game.players
	if symbol == players[0].GetSymbol() {
		return players[0].GetName()
	} else if symbol == players[1].GetSymbol() {
		return players[1].GetName()
	}
	return ""
}

func (game game) Start() string {
	fmt.Println("Lets begin!")

	grid := game.grid
	players := game.players

	grid.Print()
	for true {
		grid.Move(players[0])
		grid.Print()
		gameOver, winningSymbol := grid.Check()

		if gameOver {
			return game.winnerFromSymbol(winningSymbol)
		}

		grid.Move(players[1])
		grid.Print()
		gameOver, winningSymbol = grid.Check()

		if gameOver {
			return game.winnerFromSymbol(winningSymbol)
		}
	}
	return ""
}
