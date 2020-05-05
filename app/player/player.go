package player

type Player struct {
	symbol string
	name   string
}

func NewPlayer(symbol string, name string) Player {
	return Player{symbol, name}
}

func (p Player) GetSymbol() string {
	return p.symbol
}

func (p Player) GetName() string {
	return p.name
}
