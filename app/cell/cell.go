package cell

import "fmt"

type Cell struct {
	Value string
}

func NewCell() Cell {
	return Cell{"_"}
}

func (cell Cell) Print() {
	fmt.Printf("%v\t", cell.Value)
}

func (cell Cell) IsEmpty() bool {
	return cell.Value == "_"
}

func (cell *Cell) Assign(symbol string) {
	cell.Value = symbol
}
