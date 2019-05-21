package command

import "github.com/jnnnnn/toyrobotgolang/state"

type Move struct {
}

func (c Move) Parse(line string) bool {
	return line == "MOVE"
}

func (c Move) Execute(model state.State) {
	r := model.Robot
	if r == nil {
		return
	}
	x := r.PositionX
	y := r.PositionY
	switch r.Current {
	case state.North:
		y++
	case state.South:
		y--
	case state.East:
		x++
	case state.West:
		x--
	}
	if model.Table.ValidPosition(x, y) {
		r.PositionX = x
		r.PositionY = y
	}
}
