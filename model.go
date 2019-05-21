package main

import "errors"

// Table represents a table that the robot can drive on
type Table struct {
	SizeX, SizeY int
}

// ValidPosition is true if the given coordinates are on the table
func (t Table) ValidPosition(x, y int) bool {
	if x >= t.SizeX || x < 0 || y >= t.SizeY || y < 0 {
		return false
	}
	return true
}

// Move the robot r across the table.
func Move(r *Robot, table Table) {
	x := r.PositionX
	y := r.PositionY
	switch r.Current {
	case North:
		y++
	case South:
		y--
	case East:
		x++
	case West:
		x--
	}
	if table.ValidPosition(x, y) {
		r.PositionX = x
		r.PositionY = y
	}
}

// Place a robot on a table.
func Place(x, y int, f Facing, t Table) *Robot {
	if t.ValidPosition(x, y) {
		return &Robot{PositionX: x, PositionY: y, Current: f}
	}
	return nil
}

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

// ToString turns Facing into a string
func (f Facing) ToString() string {
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
