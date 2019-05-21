package commands

import (
	"log"

	"github.com/jnnnnn/toyrobotgolang/model"
)

// Command doesn't really belong here but it's only 3 lines of code
type Command interface {
	Parse(line string) bool
	Execute(state *model.Model)
}

// All the available commands
var All []Command

func init() {
	log.SetFlags(0)
	All = []Command{&Move{}, &Turn{}, &Report{}, &Place{}}
}
