package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/asmolik/aoc2020/internal/utils"
)

type Seat struct {
	row int64
	col int64
}

func (s *Seat) seatId() int64 {
	return s.row*8 + s.col
}

func parseInput(lines []string) []Seat {
	seats := []Seat{}
	for _, line := range lines {
		row := line[:7]
		col := line[7:]
		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")
		col = strings.ReplaceAll(col, "L", "0")
		col = strings.ReplaceAll(col, "R", "1")
		row_i, _ := strconv.ParseInt(row, 2, 64)
		col_i, _ := strconv.ParseInt(col, 2, 64)
		seats = append(seats, Seat{row_i, col_i})
	}
	return seats
}

func main() {
	data_str := utils.ReadLines("input")
	seats := parseInput(data_str)
	fmt.Println(part1(seats))
	fmt.Println(part2(seats))
}

func part1(seats []Seat) int64 {
	maxId := int64(-1)
	for _, seat := range seats {
		tmpId := seat.seatId()
		if tmpId > maxId {
			maxId = tmpId
		}
	}
	return maxId
}

func part2(seats []Seat) int64 {
	maxId := int64(-1)
	minId := int64(math.MaxInt64)
	takenSeats := make(map[int64]bool)
	for _, seat := range seats {
		tmpId := seat.seatId()
		if tmpId > maxId {
			maxId = tmpId
		}
		if tmpId < minId {
			minId = tmpId
		}
		takenSeats[tmpId] = true
	}
	myId := int64(-1)
	for seat := range takenSeats {
		if seat == minId || seat == maxId {
			continue
		}
		_, before := takenSeats[seat-1]
		_, after := takenSeats[seat+1]
		if !before {
			myId = seat - 1
			break
		}
		if !after {
			myId = seat + 1
			break
		}
	}
	return myId
}
