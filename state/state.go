package state

type State struct {
	*Robot
	Table
}

// Robot is a thing on a Table
type Robot struct {
	PositionX int
	PositionY int
	Current   Facing
}
