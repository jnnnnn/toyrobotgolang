package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jnnnnn/toyrobotgolang/model"
)

// Place: a command to put a robot at a certain position and orientation
type Place struct {
	x, y   int
	facing model.Facing
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

func (c Place) Execute(state *model.Model) {
	if state.Table.ValidPosition(c.x, c.y) {
		state.Robot = &model.Robot{PositionX: c.x, PositionY: c.y, Current: c.facing}
	}
}

// ParseFacing is a function
func ParseFacing(s string) (model.Facing, error) {
	switch s {
	case "NORTH":
		return model.Facing_North, nil
	case "EAST":
		return model.Facing_East, nil
	case "SOUTH":
		return model.Facing_South, nil
	case "WEST":
		return model.Facing_West, nil
	}
	return -1, errors.New("Unparseable facing direction")
}
