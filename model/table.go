package model

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
