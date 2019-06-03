package commands

import (
	"testing"

	"github.com/jnnnnn/toyrobotgolang/model"
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
	state := testInitialModel()
	c := Turn{model.Turning_Left}

	c.Execute(state)

	if state.Robot != nil {
		t.Errorf("Robot was not initialized but still Turnd")
	}
}

func TestTurnValid(t *testing.T) {
	state := testModel()
	c := Turn{model.Turning_Left}

	c.Execute(state)

	if state.Robot.Current != model.Facing_West {
		t.Errorf("Turn left from north is west")
	}
}
