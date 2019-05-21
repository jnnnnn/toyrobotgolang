package commands

import (
	"testing"

	"github.com/jnnnnn/toyrobotgolang/state"
)

func TestPlaceParseGood(t *testing.T) {
	c := Place{}
	if c.Parse("PLACE 0 0 NORTH") != true {
		t.Errorf("Should have parsed Place.")
	}
}
func TestPlaceParseBad(t *testing.T) {
	c := Place{}
	if c.Parse("PLACE 0 0 0") == true {
		t.Errorf("Shouldn't have parsed invalid command.")
	}
}

func TestPlaceValid(t *testing.T) {
	model := testInitialState()
	c := Place{2, 3, state.West}

	c.Execute(model)

	if model.Robot == nil {
		t.Errorf("Robot was not placed")
	}
	r := model.Robot
	if r.PositionX != 2 || r.PositionY != 3 || r.Current != state.West {
		t.Errorf("Robot was placed incorrectly.")
	}
}

func TestPlaceInvalid(t *testing.T) {
	model := &state.State{Table: state.Table{5, 5}, Robot: &state.Robot{2, 4, state.North}}
	c := Place{2, 6, state.West}

	c.Execute(model)

	r := model.Robot
	if r.PositionY != 4 && r.PositionX != 2 {
		t.Errorf("Robot Placed off table to %d %d", r.PositionX, r.PositionY)
	}
}
