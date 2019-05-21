package main

import (
	"bufio"
	"os"

	"github.com/jnnnnn/toyrobotgolang/commands"

	"github.com/jnnnnn/toyrobotgolang/state"
)

func main() {
	model := state.Initial()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, command := range commands.All {
			if command.Parse(scanner.Text()) {
				command.Execute(model)
			}
		}
	}
}
