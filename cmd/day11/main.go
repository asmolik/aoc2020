package main

import (
	"fmt"

	"github.com/asmolik/aoc2020/internal/utils"
)

const floor = 0
const free = 1
const occupied = 2

type Ferry struct {
	seats             [][]int
	adjacentSeats     Adjacent
	occupiedTolerance int
}

func (ferry *Ferry) size() (int, int) {
	return len(ferry.seats) - 1, len(ferry.seats[0]) - 1
}

type Adjacent func(int, int, int, int) [][2]int

func (ferry *Ferry) shuffle() (Ferry, bool) {
	ret := [][]int{}
	shuffled := false
	for i, row := range ferry.seats {
		newRow := []int{}
		for j, seat := range row {
			newSeat := ferry.shuffleSeat(seat, i, j)
			newRow = append(newRow, newSeat)
			if newSeat != seat {
				shuffled = true
			}
		}
		ret = append(ret, newRow)
	}
	return Ferry{ret, ferry.adjacentSeats, ferry.occupiedTolerance}, shuffled
}

func (ferry *Ferry) isSeatOccupied(row int, col int) bool {
	switch ferry.seats[row][col] {
	case floor:
		return false
	case free:
		return false
	case occupied:
		return true
	}
	return false
}

func (ferry *Ferry) shuffleSeat(seat int, row int, column int) int {
	maxRow, maxCol := ferry.size()
	adjacentSeats := ferry.adjacentSeats(row, column, maxRow, maxCol)
	switch seat {
	case floor:
		return floor
	case free:
		for _, seat := range adjacentSeats {
			if ferry.isSeatOccupied(seat[0], seat[1]) {
				return free
			}
		}
		return occupied
	case occupied:
		numOccupied := 0
		for _, seat := range adjacentSeats {
			if ferry.isSeatOccupied(seat[0], seat[1]) {
				numOccupied++
			}
		}
		if numOccupied >= ferry.occupiedTolerance {
			return free
		}
		return occupied
	}
	return floor
}

func (ferry *Ferry) occupiedSeats() int {
	sum := 0
	for _, row := range ferry.seats {
		for _, seat := range row {
			if seat == occupied {
				sum++
			}
		}
	}
	return sum
}

func parseInput(lines []string) Ferry {
	ret := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			switch char {
			case '.':
				row = append(row, floor)
			case '#':
				row = append(row, occupied)
			case 'L':
				row = append(row, free)
			}
		}
		ret = append(ret, row)
	}
	return Ferry{ret, utils.AdjacentCells, 4}
}

func main() {
	data_str := utils.ReadLines("input")
	run_for_data(data_str)
}

func run_for_data(data []string) {
	ferry := parseInput(data)
	fmt.Println(part1(ferry))
	fmt.Println(part2(ferry))
}

func part1(ferry Ferry) int {
	for {
		shuffled := false
		ferry, shuffled = ferry.shuffle()
		if !shuffled {
			break
		}
	}
	return ferry.occupiedSeats()
}

func part2(ferry Ferry) int {
	ferry.adjacentSeats = ferry.adjacent
	ferry.occupiedTolerance = 5
	for {
		shuffled := false
		ferry, shuffled = ferry.shuffle()
		if !shuffled {
			break
		}
	}
	return ferry.occupiedSeats()
}

func (ferry *Ferry) adjacent(row int, col int, maxRow int, maxCol int) [][2]int {
	ret := [][2]int{}
	directions := createDirections()
	for _, direction := range directions {
		seat := ferry.findNeighbor(direction, row, col, maxRow, maxCol)
		if seat[0] >= 0 {
			ret = append(ret, seat)
		}
	}
	return ret
}

func (ferry *Ferry) findNeighbor(direction [2]int, row int, col int, maxRow int, maxCol int) [2]int {
	ret := [2]int{-1, -1}
	i := row + direction[0]
	j := col + direction[1]
	for i >= 0 && i <= maxRow && j >= 0 && j <= maxCol {
		if ferry.seats[i][j] != floor {
			return [2]int{i, j}
		}
		i += direction[0]
		j += direction[1]
	}
	return ret
}

func createDirections() [][2]int {
	directions := [][2]int{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			directions = append(directions, [2]int{i, j})
		}
	}
	return directions
}
