package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jnnnnn/toyrobotgolang/state"
)

// Place: a command to put a robot at a certain position and orientation
type Place struct {
	x, y   int
	facing state.Facing
}

func (c *Place) Parse(line string) bool {
	words := strings.Split(line, " ")
	if len(words) != 4 {
		return false
	}
	if words[0] != "PLACE" {
		return false
	}

	var facingstr string
	_, err := fmt.Sscanf(line, "PLACE %d %d %s\n", &c.x, &c.y, &facingstr)
	if err != nil {
		fmt.Printf("Scan error %s", err.Error())
	}

	facing, err := ParseFacing(facingstr)
	if err != nil {
		return false
	}
	c.facing = facing

	return true
}

func (c Place) Execute(model *state.State) {
	if model.Table.ValidPosition(c.x, c.y) {
		model.Robot = &state.Robot{PositionX: c.x, PositionY: c.y, Current: c.facing}
	}
}

// ParseFacing is a function
func ParseFacing(s string) (state.Facing, error) {
	switch s {
	case "NORTH":
		return state.North, nil
	case "EAST":
		return state.East, nil
	case "SOUTH":
		return state.South, nil
	case "WEST":
		return state.West, nil
	}
	return -1, errors.New("Unparseable facing direction")
}
