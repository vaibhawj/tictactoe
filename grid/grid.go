package grid

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vaibhawj/tictactoe/cell"
	"github.com/vaibhawj/tictactoe/player"
)

type Grid struct {
	cells [3][3]cell.Cell
}

func NewGrid() *Grid {
	return &Grid{
		[3][3]cell.Cell{
			{cell.NewCell("_"), cell.NewCell("_"), cell.NewCell("_")},
			{cell.NewCell("_"), cell.NewCell("_"), cell.NewCell("_")},
			{cell.NewCell("_"), cell.NewCell("_"), cell.NewCell("_")}}}
}

func (grid Grid) Print() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid.cells[i][j].Print()
		}
		fmt.Println()
	}
}

func (grid Grid) validate(x int64, y int64) (status bool, msg string) {
	status = true
	if !grid.cells[x][y].IsEmpty() {
		status = false
		msg = "This place is already filled"
	}
	return
}

func (grid *Grid) assign(x int64, y int64, symbol string) {
	grid.cells[x][y].Assign(symbol)
}

func (grid Grid) Check() (gameOver bool, winner string) {
	cells := grid.cells
	gameOver = false
	winner = ""

	// rows equal
	if !cells[0][0].IsEmpty() && cells[0][0] == cells[0][1] && cells[0][1] == cells[0][2] {
		gameOver = true
		winner = cells[0][0].Value
	}

	if !cells[1][0].IsEmpty() && cells[1][0] == cells[1][1] && cells[1][1] == cells[1][2] {
		gameOver = true
		winner = cells[1][0].Value
	}

	if !cells[2][0].IsEmpty() && cells[2][0] == cells[2][1] && cells[2][1] == cells[2][2] {
		gameOver = true
		winner = cells[2][0].Value
	}

	// columns equal
	if !cells[0][0].IsEmpty() && cells[0][0] == cells[1][0] && cells[1][0] == cells[2][0] {
		gameOver = true
		winner = cells[0][0].Value
	}

	if !cells[0][1].IsEmpty() && cells[0][1] == cells[1][1] && cells[1][1] == cells[2][1] {
		gameOver = true
		winner = cells[0][1].Value
	}

	if !cells[0][2].IsEmpty() && cells[0][2] == cells[1][2] && cells[1][2] == cells[2][2] {
		gameOver = true
		winner = cells[0][2].Value
	}

	// Diagonals equal
	if !cells[0][0].IsEmpty() && cells[0][0] == cells[1][1] && cells[1][1] == cells[2][2] {
		gameOver = true
		winner = cells[0][0].Value
	}

	// Diagonals equal
	if !cells[0][2].IsEmpty() && cells[0][2] == cells[1][1] && cells[1][1] == cells[2][0] {
		gameOver = true
		winner = cells[0][2].Value
	}

	if gameOver {
		return
	}

	// check all filled
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if cells[i][j].IsEmpty() {
				return
			}
		}
	}
	gameOver = true

	return
}

func (grid *Grid) Move(p player.Player) {
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

		status, msg := grid.validate(x, y)
		if !status {
			fmt.Println(msg)
			continue
		}

		grid.assign(x, y, p.GetSymbol())
		break
	}
}
