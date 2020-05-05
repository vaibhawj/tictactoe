package cell

import "fmt"

type Cell struct {
	Value string
}

func NewCell(val string) Cell {
	return Cell{val}
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
