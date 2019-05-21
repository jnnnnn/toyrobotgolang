package main

import "testing"

func TestValidPlace(t *testing.T) {
	table := Table{SizeX: 5, SizeY: 5}
	robot := Place(0, 0, North, table)
	if robot.PositionX != 0 || robot.PositionY != 0 || robot.Current != North {
		t.Errorf("Robot not placed correctly")
	}
}

func TestInvalidPlace(t *testing.T) {
	table := Table{SizeX: 5, SizeY: 5}
	robot := Place(10, 0, North, table)
	if robot != nil {
		t.Errorf("Robot placed in an invalid position")
	}
}

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

func TestTurnLeft(t *testing.T) {
	robot := Robot{Current: North}
	robot.Turn(Left)
	if robot.Current != West {
		t.Errorf("Left of North is West, got %d", robot.Current)
	}
}

func TestTurnRight(t *testing.T) {
	robot := Robot{Current: West}
	robot.Turn(Right)
	if robot.Current != North {
		t.Errorf("Right of West is North")
	}
}
