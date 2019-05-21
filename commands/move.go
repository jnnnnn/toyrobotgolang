package commands

import "github.com/jnnnnn/toyrobotgolang/model"

type Move struct {
}

func (c Move) Parse(line string) bool {
	return line == "MOVE"
}

func (c Move) Execute(state *model.Model) {
	r := state.Robot
	if r == nil {
		return
	}
	x := r.PositionX
	y := r.PositionY
	switch r.Current {
	case model.North:
		y++
	case model.South:
		y--
	case model.East:
		x++
	case model.West:
		x--
	}
	if state.Table.ValidPosition(x, y) {
		r.PositionX = x
		r.PositionY = y
	}
}
