package commands

import (
	"log"

	"github.com/jnnnnn/toyrobotgolang/model"
)

type Report struct {
}

func (c *Report) Parse(line string) bool {
	return line == "REPORT"
}

func (c Report) Execute(state *model.Model) {
	r := state.Robot
	if r == nil {
		return
	}

	log.Printf("Robot at %d %d %s\n", r.PositionX, r.PositionY, r.Current.ToString())
}
