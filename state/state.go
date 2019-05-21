package state

type State struct {
	*Robot
	Table
}

// Table represents a table that the robot can drive on
type Table struct {
	SizeX, SizeY int
}

// ValidPosition is true if the given coordinates are on the table
func (t Table) ValidPosition(x, y int) bool {
	if x >= t.SizeX || x < 0 || y >= t.SizeY || y < 0 {
		return false
	}
	return true
}

// Facing is a compass point
type Facing int

// These are the allowed Facings in clockwise order
const (
	North Facing = iota
	East
	South
	West
	FacingCount
)

// ToString turns Facing into a string
func (f Facing) ToString() string {
	switch f {
	case North:
		return "NORTH"
	case East:
		return "EAST"
	case South:
		return "SOUTH"
	case West:
		return "WEST"
	}
	return ""
}

// Turning left or right
type Turning int

// The two ways that the robot can turn
const (
	Left Turning = iota
	Right
)

// Robot is a thing on a Table
type Robot struct {
	PositionX int
	PositionY int
	Current   Facing
}

// Turn the robot left or right
func (r *Robot) Turn(turn Turning) {
	if turn == Left {
		r.Current--
	} else {
		r.Current++
	}
	r.Current %= FacingCount
	if r.Current < 0 {
		r.Current += FacingCount
	}
}
