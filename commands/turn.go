package commands

import "github.com/jnnnnn/toyrobotgolang/model"

type Turn struct {
	model.Turning
}

func (c *Turn) Parse(line string) bool {
	switch line {
	case "LEFT":
		c.Turning = model.Turning_Left
		return true
	case "RIGHT":
		c.Turning = model.Turning_Right
		return true
	default:
		return false
	}
}

func (c Turn) Execute(state *model.Model) {
	r := state.Robot
	if r == nil {
		return
	}
	if c.Turning == model.Turning_Left {
		r.Current--
	} else {
		r.Current++
	}
	r.Current %= model.Facing_Count
	if r.Current < 0 {
		r.Current += model.Facing_Count
	}
}
