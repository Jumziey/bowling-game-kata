package bowling_test

import (
	"testing"

	"example.com/bowling"
	"github.com/stretchr/testify/assert"
)

func TestGutterGame(t *testing.T) {
	g := bowling.NewGame()

	for _, pins := range make([]uint, 20, 20) {
		g.Roll(pins)
	}
	assert.Zero(t, g.Score())
}
