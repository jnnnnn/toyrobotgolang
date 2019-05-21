package commands

import (
	"testing"

	"github.com/jnnnnn/toyrobotgolang/state"
)

func TestTurnParseGood(t *testing.T) {
	c := Turn{}
	if c.Parse("LEFT") != true {
		t.Errorf("Should have parsed LEFT.")
	}
}
func TestTurnParseBad(t *testing.T) {
	c := Turn{}
	if c.Parse("RIGHTR") == true {
		t.Errorf("Shouldn't have parsed RIGHTR.")
	}
}

func TestTurnInitialize(t *testing.T) {
	model := testInitialState()
	c := Turn{state.Left}

	c.Execute(model)

	if model.Robot != nil {
		t.Errorf("Robot was not initialized but still Turnd")
	}
}

func TestTurnValid(t *testing.T) {
	model := testState()
	c := Turn{state.Left}

	c.Execute(model)

	if model.Robot.Current != state.West {
		t.Errorf("Turn left from north is west")
	}
}
