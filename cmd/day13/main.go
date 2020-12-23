package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/asmolik/aoc2020/internal/utils"
)

func parseInput(lines []string) (int, []int) {
	ret := []int{}
	minTimestamp, _ := strconv.Atoi(lines[0])
	s := strings.Split(lines[1], ",")
	for _, i := range s {
		if i != "x" {
			tmp, _ := strconv.Atoi(i)
			ret = append(ret, tmp)
		} else {
			ret = append(ret, 0)
		}
	}
	return minTimestamp, ret
}

func main() {
	data_str := utils.ReadLines("input")
	run_for_data(data_str)
}

func run_for_data(data []string) {
	minTimestamp, schedules := parseInput(data)
	fmt.Println(part1(minTimestamp, schedules))
	fmt.Println(part2(schedules))
}

func part1(minTimestamp int, schedules []int) int {
	waitTime := math.MaxInt64
	busId := -1
	for _, i := range schedules {
		if i == 0 {
			continue
		}
		tmp := minTimestamp / i
		tmp = tmp * i
		if tmp == minTimestamp {
			return 0
		}
		tmp += i
		wait := tmp - minTimestamp
		if wait < waitTime {
			busId = i
			waitTime = wait
		}
	}
	return waitTime * busId
}

func part2(schedules []int) int {
	step := schedules[0]
	timestamp := 0
	for i, busID := range schedules[0:] {
		if i == 0 || busID == 0 {
			continue
		}
		for (timestamp+i)%busID != 0 {
			timestamp += step
		}
		step *= busID
	}
	return timestamp
}
