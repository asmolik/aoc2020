package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Map struct {
	grid  [][]rune
	x, y  int
	trees int
}

type Move struct {
	x, y int
}

func (m *Map) move(v Move) bool {
	m.x = (m.x + v.x) % len(m.grid[0])
	m.y += v.y
	if m.y >= len(m.grid) {
		return true
	}
	if m.grid[m.y][m.x] == '#' {
		m.trees++
	}
	return false
}

func (m *Map) reset() {
	m.x = 0
	m.y = 0
	m.trees = 0
}

func readInput(lines []string) Map {
	m := Map{}
	for _, line := range lines {
		m.grid = append(m.grid, []rune(line))
	}
	return m
}

func main() {
	test_data, _ := ioutil.ReadFile("test_input")
	test_data_str := strings.Split(string(test_data), "\n")
	test_map := readInput(test_data_str)
	fmt.Println(part1(test_map, Move{3, 1}))
	test_map.reset()
	fmt.Println(part2(test_map))
	data, _ := ioutil.ReadFile("input")
	data_str := strings.Split(string(data), "\n")
	trip_map := readInput(data_str)
	fmt.Println(part1(trip_map, Move{3, 1}))
	trip_map.reset()
	fmt.Println(part2(trip_map))
}

func part1(m Map, v Move) int {
	arrived := false
	for !arrived {
		arrived = m.move(v)
	}
	return m.trees
}

func part2(m Map) int {
	moves := []Move{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	trees := []int{}
	for _, move := range moves {
		trees = append(trees, part1(m, move))
		m.reset()
	}
	product := 1
	for _, i := range trees {
		product *= i
	}
	return product
}
