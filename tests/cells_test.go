package tests

import (
	"testing"

	"github.com/vaibhawj/tictactoe/app/cell"
)

func TestIsEmpty(t *testing.T) {
	c := cell.NewCell()
	if !c.IsEmpty() {
		t.Error("New cell is not empty i.e. \"_\"")
	}
}

func TestAssign(t *testing.T) {
	c := cell.NewCell()
	c.Assign("X")
	if c.Value != "X" {
		t.Errorf("Expected Cell value after assign is X but got %v", c.Value)
	}
}
