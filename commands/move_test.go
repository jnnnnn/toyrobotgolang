package commands

import (
	"testing"

	"github.com/jnnnnn/toyrobotgolang/model"
)

func testModel() *model.Model {
	return &model.Model{Table: model.Table{5, 5}, Robot: &model.Robot{2, 2, model.North}}
}

func testInitialModel() *model.Model {
	return &model.Model{Table: model.Table{5, 5}, Robot: nil}
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
	state := testInitialModel()
	c := Move{}

	c.Execute(state)

	if state.Robot != nil {
		t.Errorf("Robot was not initialized but still moved")
	}
}

func TestMoveValid(t *testing.T) {
	state := testModel()
	c := Move{}

	c.Execute(state)

	r := state.Robot
	if r.PositionY != 3 && r.PositionX != 2 {
		t.Errorf("Move to 3,2 didn't work; got %d %d", r.PositionX, r.PositionY)
	}
}
func TestMoveInvalid(t *testing.T) {
	state := &model.Model{Table: model.Table{5, 5}, Robot: &model.Robot{2, 4, model.North}}
	c := Move{}

	c.Execute(state)

	r := state.Robot
	if r.PositionY != 4 && r.PositionX != 2 {
		t.Errorf("Robot moved off table to %d %d", r.PositionX, r.PositionY)
	}
}
