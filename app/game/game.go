package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vaibhawj/tictactoe/app/grid"
	"github.com/vaibhawj/tictactoe/app/player"
)

func InitializePlayer(i int, symbol string) player.Player {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter name of player %v\n", i)
	input, _ := reader.ReadString('\n')
	name := strings.Replace(input, "\n", "", 1)
	return player.NewPlayer(symbol, name)
}

type game struct {
	grid    *grid.Grid
	players [2]player.Player
}

func NewGame(players [2]player.Player) *game {
	return &game{
		grid.NewGrid(),
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
		game.move(players[0])
		grid.Print()
		gameOver, winningSymbol := grid.Check()

		if gameOver {
			return game.winnerFromSymbol(winningSymbol)
		}

		game.move(players[1])
		grid.Print()
		gameOver, winningSymbol = grid.Check()

		if gameOver {
			return game.winnerFromSymbol(winningSymbol)
		}
	}
	return ""
}

func (game *game) move(p player.Player) {
	grid := game.grid
	for true {
		fmt.Printf("%v's turn (%v). Enter coordinates eg. 00, 01 or 22\n", p.GetName(), p.GetSymbol())
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')

		x, err := strconv.ParseInt(strings.Split(choice, "")[0], 10, 8)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		y, err := strconv.ParseInt(strings.Split(choice, "")[1], 10, 8)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		if (x < 0 || x > 2) || (y < 0 || y > 2) {
			fmt.Println("Invalid input")
			continue
		}

		if !grid.IsEmpty(x, y) {
			fmt.Println("This place is already filled")
			continue
		}

		grid.Assign(x, y, p.GetSymbol())
		break
	}
}
