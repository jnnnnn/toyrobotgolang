package main

import "errors"

// Facing is a Facing point
type Facing int

// These are the allowed Facings in clockwise order
const (
	North Facing = iota
	East
	South
	West
	FacingCount
)

func UnParseFacing(f Facing) string {
	switch f {
	case North:
		return "NORTH"
	case East:
		return "EAST"
	case South:
		return "SOUTH"
	case West:
		return "WEST"
	}
	return ""
}

// ParseFacing is a function
func ParseFacing(s string) (Facing, error) {
	switch s {
	case "NORTH":
		return North, nil
	case "EAST":
		return East, nil
	case "SOUTH":
		return South, nil
	case "WEST":
		return West, nil
	}
	return -1, errors.New("Unparseable facing direction")
}

// Turning left or right
type Turning int

// The two ways that the robot can turn
const (
	Left Turning = iota
	Right
)

// Robot is a thing on a Table
type Robot struct {
	PositionX int
	PositionY int
	Current   Facing
}

// Turn the robot left or right
func (r *Robot) Turn(turn Turning) {
	if turn == Left {
		r.Current--
	} else {
		r.Current++
	}
	r.Current %= FacingCount
	if r.Current < 0 {
		r.Current += FacingCount
	}
}
