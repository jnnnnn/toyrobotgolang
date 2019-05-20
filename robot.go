package main

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
