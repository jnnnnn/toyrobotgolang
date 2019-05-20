package main

import "testing"

func TestMoveValid(t *testing.T) {
	table := Table{SizeX: 5, SizeY: 5}
	robot := Robot{PositionX: 2, PositionY: 2, Current: North}
	Move(robot, table)
	if robot.PositionY != 3 && robot.PositionX != 2 {
		t.Errorf("Move to 3,2 didn't work; got %d %d", robot.PositionX, robot.PositionY)
	}
}

func TestMoveInvalid(t *testing.T) {
	table := Table{SizeX: 5, SizeY: 5}
	robot := Robot{PositionX: 4, PositionY: 2, Current: East}
	Move(robot, table)
	if robot.PositionY != 2 && robot.PositionX != 4 {
		t.Errorf("Robot moved off table to %d %d", robot.PositionX, robot.PositionY)
	}
}

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
