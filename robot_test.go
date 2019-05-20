package main

import "testing"

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
