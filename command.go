package main

// Command doesn't really belong here but it's only 3 lines of code
type Command interface {
	Parse(line string) bool
	Execute(state State)
}
