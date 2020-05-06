package grid

import (
	"fmt"

	"github.com/vaibhawj/tictactoe/app/cell"
)

type Grid struct {
	cells [3][3]cell.Cell
}

func NewGrid() *Grid {
	return &Grid{
		[3][3]cell.Cell{
			{cell.NewCell(), cell.NewCell(), cell.NewCell()},
			{cell.NewCell(), cell.NewCell(), cell.NewCell()},
			{cell.NewCell(), cell.NewCell(), cell.NewCell()}}}
}

func (grid Grid) Print() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid.cells[i][j].Print()
		}
		fmt.Println()
	}
}

func (grid Grid) IsEmpty(x int64, y int64) bool {
	return grid.cells[x][y].IsEmpty()
}

func (grid *Grid) Assign(x int64, y int64, symbol string) {
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
