package model

type Model struct {
	*Robot
	Table
}

// Robot is a thing on a Table
type Robot struct {
	PositionX int
	PositionY int
	Current   Facing
}

func Initial() *Model {
	return &Model{Table: Table{SizeX: 5, SizeY: 5}}
}
