package commands

import (
	"testing"

	"github.com/jnnnnn/toyrobotgolang/model"
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
	state := testInitialModel()
	c := Place{2, 3, model.West}

	c.Execute(state)

	if state.Robot == nil {
		t.Errorf("Robot was not placed")
	}
	r := state.Robot
	if r.PositionX != 2 || r.PositionY != 3 || r.Current != model.West {
		t.Errorf("Robot was placed incorrectly.")
	}
}

func TestPlaceInvalid(t *testing.T) {
	state := &model.Model{Table: model.Table{5, 5}, Robot: &model.Robot{2, 4, model.North}}
	c := Place{2, 6, model.West}

	c.Execute(state)

	r := state.Robot
	if r.PositionY != 4 && r.PositionX != 2 {
		t.Errorf("Robot Placed off table to %d %d", r.PositionX, r.PositionY)
	}
}
