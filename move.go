package main

type Move struct {
}

func (c Move) Parse(line string) bool {
	return line == "MOVE"
}

func (c Move) Execute(state State) {
	r := state.Robot
	if r == nil {
		return
	}
	x := r.PositionX
	y := r.PositionY
	switch r.Current {
	case North:
		y++
	case South:
		y--
	case East:
		x++
	case West:
		x--
	}
	if state.Table.ValidPosition(x, y) {
		r.PositionX = x
		r.PositionY = y
	}
}
