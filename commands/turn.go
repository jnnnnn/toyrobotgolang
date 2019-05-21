package commands

import "github.com/jnnnnn/toyrobotgolang/state"

type Turn struct {
	state.Turning
}

func (c *Turn) Parse(line string) bool {
	switch line {
	case "LEFT":
		c.Turning = state.Left
		return true
	case "RIGHT":
		c.Turning = state.Right
		return true
	default:
		return false
	}
}

func (c Turn) Execute(model *state.State) {
	r := model.Robot
	if r == nil {
		return
	}
	if c.Turning == state.Left {
		r.Current--
	} else {
		r.Current++
	}
	r.Current %= state.FacingCount
	if r.Current < 0 {
		r.Current += state.FacingCount
	}
}
