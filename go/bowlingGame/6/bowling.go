package bowling

type Game struct{}

func NewGame() Game { return Game{} }

func (g *Game) Roll(pins uint) {}

func (g *Game) Score() uint {
	return 0
}
