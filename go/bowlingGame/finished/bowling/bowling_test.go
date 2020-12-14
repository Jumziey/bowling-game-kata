package bowling_test

import (
	"testing"

	"example.com/bowling"
	"github.com/stretchr/testify/assert"
)

func rollMany(g *bowling.Game, nrOfRolls, pins uint) {
	for i := uint(0); i < nrOfRolls; i++ {
		g.Roll(pins)
	}
}

func rollSpare(g *bowling.Game) {
	g.Roll(5)
	g.Roll(5)
}

func rollPerfectSpare(g *bowling.Game) {
	g.Roll(9)
	g.Roll(1)
}

func rollStrike(g *bowling.Game) {
	g.Roll(10)
}

func rollFrame(g *bowling.Game, first, sec uint) {
	g.Roll(first)
	if first < 10 {
		g.Roll(sec)
	}
}

func TestGutterGame(t *testing.T) {
	g := bowling.NewGame()

	rollMany(&g, 20, 0)

	assert.Zero(t, g.Score())
}

func TestAllOnes(t *testing.T) {
	g := bowling.NewGame()

	rollMany(&g, 20, 1)

	assert.Equal(t, 20, int(g.Score()))
}

func TestOneSpar(t *testing.T) {
	g := bowling.NewGame()

	rollSpare(&g)
	g.Roll(3)
	rollMany(&g, 17, 0)

	assert.Equal(t, 5+5+3+3, int(g.Score()))
}

func TestNoSpareBetweenFrames(t *testing.T) {
	g := bowling.NewGame()

	rollFrame(&g, 3, 5)
	rollFrame(&g, 5, 3)
	rollMany(&g, 18, 0)

	assert.Equal(t, 3+5+5+3, int(g.Score()))
}

func TestBestSpareGame(t *testing.T) {
	g := bowling.NewGame()

	for i := 0; i < 10; i++ {
		rollPerfectSpare(&g)
	}
	g.Roll(9)

	assert.Equal(t, 190, int(g.Score()))
}

func TestOneStrike(t *testing.T) {
	g := bowling.NewGame()

	rollStrike(&g)
	g.Roll(3)
	g.Roll(7)
	rollMany(&g, 16, 0)

	assert.Equal(t, 10+(3+7)*2, int(g.Score()))
}

func TestPerfectGame(t *testing.T) {
	g := bowling.NewGame()

	for i := 0; i < 13; i++ {
		rollStrike(&g)
	}

	assert.Equal(t, 300, int(g.Score()))
}
