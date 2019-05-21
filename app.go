package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var robot *Robot
	table := Table{SizeX: 5, SizeY: 5}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Handle(scanner.Text(), &robot, table)
	}
}

// Handle a command for the robot
func Handle(line string, robot **Robot, table Table) {
	words := strings.Split(line, " ")
	if len(words) < 1 {
		return
	}
	if words[0] == "PLACE" {
		fmt.Println("PLACE...")
		var x, y int
		var facingstr string
		_, err := fmt.Sscanf(line, "PLACE %d %d %s\n", &x, &y, &facingstr)
		if err != nil {
			fmt.Printf("Scan error %s", err.Error())
		}
		facing, err := ParseFacing(facingstr)
		if err != nil {
			return
		}

		fmt.Println("Place()...")
		newrobot := Place(x, y, facing, table)
		if newrobot != nil {
			*robot = newrobot
		}
	}
	if *robot == nil {
		return
	}
	r := *robot
	switch words[0] {
	case "LEFT":
		r.Turn(Left)
	case "RIGHT":
		r.Turn(Right)
	case "MOVE":
		Move(r, table)
	case "REPORT":
		fmt.Printf("r is at %d %d %s\n", r.PositionX, r.PositionY, r.Current.ToString())
	}
}
