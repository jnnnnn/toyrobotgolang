package state

import "testing"

func TestTableValid(t *testing.T) {
	table := Table{SizeX: 5, SizeY: 5}
	got := table.ValidPosition(0, 4)
	if got != true {
		t.Errorf("Position 2,2 should be valid")
	}
}

func TestTablePositionLarge(t *testing.T) {
	table := Table{SizeX: 5, SizeY: 5}
	if table.ValidPosition(-1, 3) != false {
		t.Errorf("Position -1, 3 is outside bounds")
	}
}
