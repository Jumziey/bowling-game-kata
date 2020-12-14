package bowling

type Game struct {
	rolls []uint
}

func (g *Game) frames() []frame {
	frames := []frame{}
	f := frame{game: g}

	for i, _ := range g.rolls {
		f.rollIndex = append(f.rollIndex, uint(i))
		if len(f.rollIndex) >= 2 || f.game.rolls[f.rollIndex[0]] == 10 {
			frames = append(frames, f)
			if len(frames) == 10 {
				//Only have 10 frames in bowling
				//any extra is for strike/spare on
				//the last frame
				return frames
			}
			f = frame{game: g}
		}
	}

	return frames
}

func NewGame() Game {
	return Game{}
}

func (g *Game) Roll(pins uint) {
	g.rolls = append(g.rolls, pins)
}

func (g *Game) Score() uint {
	frames := g.frames()
	score := uint(0)
	for _, f := range frames {
		score += f.score()
	}
	return score
}

type frame struct {
	rollIndex []uint
	game      *Game
}

func (f *frame) rolls() []uint {
	rolls := []uint{}
	for _, index := range f.rollIndex {
		rolls = append(rolls, f.game.rolls[index])
	}
	return rolls
}

func (f *frame) lastRollIndex() uint {
	if len(f.rollIndex) == 2 {
		return f.rollIndex[1]
	} else {
		return f.rollIndex[0]
	}
}

func (f *frame) isStrike() bool {
	return f.rolls()[0] == 10
}

func (f *frame) isSpare() bool {
	rolls := f.rolls()
	return len(rolls) > 1 && rolls[0]+rolls[1] == 10
}

func (f *frame) scoreNormal() uint {
	rolls := f.rolls()
	if len(rolls) < 2 {
		return rolls[0]
	}
	return rolls[0] + rolls[1]
}

//Spare calculates on the next roll, not frame
func (f *frame) scoreSpare() uint {
	return f.scoreNormal() + f.game.rolls[f.rollIndex[1]+1]
}

//Strike calculates on the next two rolls, not frame
func (f *frame) scoreStrike() uint {
	return f.scoreNormal() +
		f.game.rolls[f.lastRollIndex()+1] +
		f.game.rolls[f.lastRollIndex()+2]
}

func (f *frame) score() uint {
	if f.isSpare() {
		return f.scoreSpare()
	} else if f.isStrike() {
		return f.scoreStrike()
	} else {
		return f.scoreNormal()
	}
}
