package bowling_test

import (
	"testing"

	"example.com/bowling"
)

func TestGutterGame(t *testing.T) {
	g := bowling.NewGame()

	for _, pins := range make([]uint, 20, 20) {
		g.Roll(pins)
	}
}
