package model

// Facing is a compass point
type Facing int

// These are the allowed Facings in clockwise order
const (
	Facing_North Facing = iota
	Facing_East
	Facing_South
	Facing_West
	Facing_Count
)

// ToString turns Facing into a string
func (f Facing) ToString() string {
	switch f {
	case Facing_North:
		return "NORTH"
	case Facing_East:
		return "EAST"
	case Facing_South:
		return "SOUTH"
	case Facing_West:
		return "WEST"
	}
	return ""
}

// Turning left or right
type Turning int

// The two ways that the robot can turn
const (
	Turning_Left Turning = iota
	Turning_Right
)
