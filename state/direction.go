
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
