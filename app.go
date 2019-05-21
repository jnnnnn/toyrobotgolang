package main

import (
	"bufio"
	"os"

	"github.com/jnnnnn/toyrobotgolang/commands"

	"github.com/jnnnnn/toyrobotgolang/model"
)

func main() {
	state := model.Initial()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, command := range commands.All {
			if command.Parse(scanner.Text()) {
				command.Execute(state)
			}
		}
	}
}
