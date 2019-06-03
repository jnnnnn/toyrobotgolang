package commands

import "github.com/jnnnnn/toyrobotgolang/model"

type Turn struct {
	model.Turning
}

func (turn *Turn) Parse(line string) bool {
	switch line {
	case "LEFT":
		turn.Turning = model.Turning_Left
		return true
	case "RIGHT":
		turn.Turning = model.Turning_Right
		return true
	default:
		return false
	}
}

func (turn Turn) Execute(state *model.Model) {
	robot := state.Robot
	if robot == nil {
		return
	}
	if turn.Turning == model.Turning_Left {
		robot.Current--
	} else {
		robot.Current++
	}
	robot.Current %= model.Facing_Count
	if robot.Current < 0 {
		robot.Current += model.Facing_Count
	}
}
