package main

type Table struct {
	SizeX, SizeY int
}

func (t Table) ValidPosition(x, y int) bool {
	if x >= t.SizeX || x < 0 || y >= t.SizeY || y < 0 {
		return false
	}
	return true
}
