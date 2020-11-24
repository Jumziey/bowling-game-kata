package bowling

type Game struct{}

func NewGame() Game { return Game{} }

func (g *Game) Roll(pins uint) {}
