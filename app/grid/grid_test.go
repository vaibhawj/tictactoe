package grid_test

import (
	"testing"

	"github.com/vaibhawj/tictactoe/app/grid"
)

func TestCheckEmptyGrid(t *testing.T) {
	grid := grid.NewGrid()
	gameOver, winner := grid.Check()

	if gameOver {
		t.Error("Empty grid, should not be gameOver")
	}

	if winner != "" {
		t.Error("Empty grid, should not have winner")
	}
}

func TestCheckFirstRowSame(t *testing.T) {
	grid := grid.NewGrid()
	grid.Assign(0, 0, "X")
	grid.Assign(0, 1, "X")
	grid.Assign(0, 2, "X")

	gameOver, winner := grid.Check()

	if !gameOver {
		t.Error("First row should be all X and hence game should be over")
	}

	if winner != "X" {
		t.Errorf("Winner should be %v but got %v", "X", winner)
	}
}

func TestCheckSecondRowSame(t *testing.T) {
	grid := grid.NewGrid()
	grid.Assign(1, 0, "X")
	grid.Assign(1, 1, "X")
	grid.Assign(1, 2, "X")

	gameOver, winner := grid.Check()

	if !gameOver {
		t.Error("Second row should be all X and hence game should be over")
	}

	if winner != "X" {
		t.Errorf("Winner should be %v but got %v", "X", winner)
	}
}

func TestCheckThirdRowSame(t *testing.T) {
	grid := grid.NewGrid()
	grid.Assign(2, 0, "X")
	grid.Assign(2, 1, "X")
	grid.Assign(2, 2, "X")

	gameOver, winner := grid.Check()

	if !gameOver {
		t.Error("Third row should be all X and hence game should be over")
	}

	if winner != "X" {
		t.Errorf("Winner should be %v but got %v", "X", winner)
	}
}

func TestCheckFirstColumnSame(t *testing.T) {
	grid := grid.NewGrid()
	grid.Assign(0, 0, "X")
	grid.Assign(1, 0, "X")
	grid.Assign(2, 0, "X")

	gameOver, winner := grid.Check()

	if !gameOver {
		t.Error("First column should be all X and hence game should be over")
	}

	if winner != "X" {
		t.Errorf("Winner should be %v but got %v", "X", winner)
	}
}

func TestCheckSecondColumnSame(t *testing.T) {
	grid := grid.NewGrid()
	grid.Assign(0, 1, "X")
	grid.Assign(1, 1, "X")
	grid.Assign(2, 1, "X")

	gameOver, winner := grid.Check()

	if !gameOver {
		t.Error("Second column should be all X and hence game should be over")
	}

	if winner != "X" {
		t.Errorf("Winner should be %v but got %v", "X", winner)
	}
}

func TestCheckThirdColumnSame(t *testing.T) {
	grid := grid.NewGrid()
	grid.Assign(0, 2, "X")
	grid.Assign(1, 2, "X")
	grid.Assign(2, 2, "X")

	gameOver, winner := grid.Check()

	if !gameOver {
		t.Error("Third column should be all X and hence game should be over")
	}

	if winner != "X" {
		t.Errorf("Winner should be %v but got %v", "X", winner)
	}
}

func TestCheckDiagonal1Same(t *testing.T) {
	grid := grid.NewGrid()
	grid.Assign(0, 0, "O")
	grid.Assign(1, 1, "O")
	grid.Assign(2, 2, "O")

	gameOver, winner := grid.Check()

	if !gameOver {
		t.Error("Third column should be all X and hence game should be over")
	}

	if winner != "O" {
		t.Errorf("Winner should be %v but got %v", "O", winner)
	}
}

func TestCheckDiagonal2Same(t *testing.T) {
	grid := grid.NewGrid()
	grid.Assign(0, 2, "O")
	grid.Assign(1, 1, "O")
	grid.Assign(2, 0, "O")

	gameOver, winner := grid.Check()

	if !gameOver {
		t.Error("Third column should be all X and hence game should be over")
	}

	if winner != "O" {
		t.Errorf("Winner should be %v but got %v", "O", winner)
	}
}

func TestCheckAllFilled(t *testing.T) {
	grid := grid.NewGrid()
	// X	X	O
	// O	O	X
	// X	O	X
	grid.Assign(0, 0, "X")
	grid.Assign(0, 1, "X")
	grid.Assign(0, 2, "O")

	grid.Assign(1, 0, "O")
	grid.Assign(1, 1, "O")
	grid.Assign(1, 2, "X")

	grid.Assign(2, 0, "X")
	grid.Assign(2, 1, "O")
	grid.Assign(2, 2, "X")

	gameOver, winner := grid.Check()

	if !gameOver {
		t.Error("All places filled, game should be over")
	}

	if winner != "" {
		t.Error("There should not be a winner")
	}
}

func TestIsEmpty(t *testing.T) {
	grid := grid.NewGrid()

	if !grid.IsEmpty(0, 0) {
		t.Error("00 should be empty")
	}

	grid.Assign(0, 0, "X")
	if grid.IsEmpty(0, 0) {
		t.Error("00 should not be empty")
	}
}
