package main

// Model : the state of a running robot-on-table simulation
type Model struct {
	robot *Robot
	table Table
}

// Move the r across the t.
func Move(r Robot, t Table) {
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
	if t.ValidPosition(x, y) {
		r.PositionX = x
		r.PositionY = y
	}
}

// Place a robot on a table.
func Place(x, y int, f Facing, t Table) *Robot {
	if t.ValidPosition(x, y) {
		return &Robot{PositionX: x, PositionY: y, Current: f}
	}
	return nil
}
