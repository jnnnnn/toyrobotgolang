package command

import "github.com/jnnnnn/toyrobotgolang/state"

// Command doesn't really belong here but it's only 3 lines of code
type Command interface {
	Parse(line string) bool
	Execute(model *state.State)
}
