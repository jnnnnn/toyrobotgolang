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
	case model.Facing_North:
		y++
	case model.Facing_South:
		y--
	case model.Facing_East:
		x++
	case model.Facing_West:
		x--
	}
	if state.Table.ValidPosition(x, y) {
		r.PositionX = x
		r.PositionY = y
	}
}
