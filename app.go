package main

import (
	"bufio"
	"os"

	"github.com/jnnnnn/toyrobotgolang/commands"
	"github.com/jnnnnn/toyrobotgolang/state"
)

func main() {
	model := &state.State{Table: state.Table{SizeX: 5, SizeY: 5}}

	var cs = []commands.Command{&commands.Move{}, &commands.Turn{}, &commands.Report{}, &commands.Place{}}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, command := range cs {
			if command.Parse(scanner.Text()) {
				command.Execute(model)
			}
		}
	}
}
