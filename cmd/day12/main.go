package main

import (
	"fmt"
	"strconv"

	"github.com/asmolik/aoc2020/internal/utils"
)

type Moveable interface {
}

type Ship struct {
	northSouth int
	eastWest   int
	direction  int
	// 0 - north, 90 - east
}

type Waypoint struct {
	northSouth int
	eastWest   int
}

func (ship *Ship) manhattanDistance() int {
	northSouth := ship.northSouth
	if ship.northSouth < 0 {
		northSouth *= -1
	}
	eastWest := ship.eastWest
	if ship.eastWest < 0 {
		eastWest *= -1
	}
	return northSouth + eastWest
}

type Instruction interface {
	move(*Ship)
	moveWaypoint(*Waypoint, *Ship)
}

type Rotation struct {
	direction int
}

type MoveForward struct {
	distance int
}

type MoveDirectionaly struct {
	distance  int
	direction int
}

func (r Rotation) move(ship *Ship) {
	direction := (ship.direction + r.direction) % 360
	if direction < 0 {
		direction = 360 + direction
	}
	ship.direction = direction
}

func (m MoveForward) move(ship *Ship) {
	switch ship.direction {
	case 0:
		ship.northSouth += m.distance
	case 90:
		ship.eastWest += m.distance
	case 180:
		ship.northSouth -= m.distance
	case 270:
		ship.eastWest -= m.distance
	}
}

func (m MoveDirectionaly) move(ship *Ship) {
	switch m.direction {
	case 0:
		ship.northSouth += m.distance
	case 90:
		ship.eastWest += m.distance
	case 180:
		ship.northSouth -= m.distance
	case 270:
		ship.eastWest -= m.distance
	}
}

func (r Rotation) moveWaypoint(waypoint *Waypoint, ship *Ship) {
	switch r.direction {
	case 90, -270:
		tmp := waypoint.eastWest
		waypoint.eastWest = waypoint.northSouth
		waypoint.northSouth = tmp * -1
	case 180, -180:
		waypoint.northSouth *= -1
		waypoint.eastWest *= -1
	case 270, -90:
		tmp := waypoint.northSouth
		waypoint.northSouth = waypoint.eastWest
		waypoint.eastWest = tmp * -1
	}
}

func (m MoveForward) moveWaypoint(waypoint *Waypoint, ship *Ship) {
	ship.northSouth += m.distance * waypoint.northSouth
	ship.eastWest += m.distance * waypoint.eastWest
}

func (m MoveDirectionaly) moveWaypoint(waypoint *Waypoint, ship *Ship) {
	switch m.direction {
	case 0:
		waypoint.northSouth += m.distance
	case 90, -270:
		waypoint.eastWest += m.distance
	case 180, -180:
		waypoint.northSouth -= m.distance
	case 270, -90:
		waypoint.eastWest -= m.distance
	}
}

func parseInput(lines []string) []Instruction {
	ret := []Instruction{}
	for _, line := range lines {
		value, _ := strconv.Atoi(line[1:])
		switch line[0] {
		case 'L':
			ret = append(ret, Rotation{-value})
		case 'R':
			ret = append(ret, Rotation{value})
		case 'F':
			ret = append(ret, MoveForward{value})
		case 'N':
			ret = append(ret, MoveDirectionaly{value, 0})
		case 'E':
			ret = append(ret, MoveDirectionaly{value, 90})
		case 'S':
			ret = append(ret, MoveDirectionaly{value, 180})
		case 'W':
			ret = append(ret, MoveDirectionaly{value, 270})
		}
	}
	return ret
}

func main() {
	data_str := utils.ReadLines("input")
	run_for_data(data_str)
}

func run_for_data(data []string) {
	instructions := parseInput(data)
	fmt.Println(part1(Ship{0, 0, 90}, instructions))
	fmt.Println(part2(Ship{0, 0, 90}, Waypoint{1, 10}, instructions))
}

func part1(ship Ship, instructions []Instruction) int {
	for _, i := range instructions {
		i.move(&ship)
	}
	return ship.manhattanDistance()
}

func part2(ship Ship, waypoint Waypoint, instructions []Instruction) int {
	for _, i := range instructions {
		i.moveWaypoint(&waypoint, &ship)
	}
	return ship.manhattanDistance()
}
