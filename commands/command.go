package commands

import (
	"log"

	"github.com/jnnnnn/toyrobotgolang/state"
)

// Command doesn't really belong here but it's only 3 lines of code
type Command interface {
	Parse(line string) bool
	Execute(model *state.State)
}

func init() {
	log.SetFlags(0)
}
