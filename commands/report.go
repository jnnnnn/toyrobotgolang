package command

import (
	"log"

	"github.com/jnnnnn/toyrobotgolang/state"
)

type Report struct {
}

func (c *Report) Parse(line string) bool {
	return line == "REPORT"
}

func (c Report) Execute(model *state.State) {
	r := model.Robot
	if r == nil {
		return
	}

	log.Printf("Robot at %d %d %s\n", r.PositionX, r.PositionY, r.Current.ToString())
}
