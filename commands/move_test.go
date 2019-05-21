package command

import (
	"testing"

	"github.com/jnnnnn/toyrobotgolang/state"
)

func testState() *state.State {
	return &state.State{Table: state.Table{5, 5}, Robot: &state.Robot{2, 2, state.North}}
}

func testInitialState() *state.State {
	return &state.State{Table: state.Table{5, 5}, Robot: nil}
}

func TestMoveParseGood(t *testing.T) {
	c := Move{}
	if c.Parse("MOVE") != true {
		t.Errorf("Should have parsed MOVE.")
	}
}
func TestMoveParseBad(t *testing.T) {
	c := Move{}
	if c.Parse("MOVER") == true {
		t.Errorf("Shouldn't have parsed MOVER.")
	}
}

func TestMoveInitialize(t *testing.T) {
	state := testInitialState()
	c := Move{}

	c.Execute(state)

	if state.Robot != nil {
		t.Errorf("Robot was not initialized but still moved")
	}
}

func TestMoveValid(t *testing.T) {
	state := testState()
	c := Move{}

	c.Execute(state)

	r := state.Robot
	if r.PositionY != 3 && r.PositionX != 2 {
		t.Errorf("Move to 3,2 didn't work; got %d %d", r.PositionX, r.PositionY)
	}
}
func TestMoveInvalid(t *testing.T) {
	state := &state.State{Table: state.Table{5, 5}, Robot: &state.Robot{2, 4, state.North}}
	c := Move{}

	c.Execute(state)

	r := state.Robot
	if r.PositionY != 4 && r.PositionX != 2 {
		t.Errorf("Robot moved off table to %d %d", r.PositionX, r.PositionY)
	}
}
