package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/asmolik/aoc2020/internal/utils"
)

func parseInput(lines []string) []int {
	ret := []int{}
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		ret = append(ret, i)
	}
	return ret
}

func main() {
	data_str := utils.ReadLines("input")
	run_for_data(data_str, 25)
}

func run_for_data(data []string, preamble int) {
	joltages := parseInput(data)
	fmt.Println(part1(joltages))
	fmt.Println(part2(0, joltages, map[int]int{}))
}

func builtInJoltage(joltages []int) int {
	max := -1
	for _, v := range joltages {
		if v > max {
			max = v
		}
	}
	return max + 3
}

func joltageDistribution(joltages []int) map[int]int {
	sort.Ints(joltages)
	previous := joltages[0]
	distribution := map[int]int{}
	for _, current := range joltages[1:] {
		distribution[current-previous] += 1
		previous = current
	}
	distribution[1] += 1
	distribution[3] += 1
	return distribution
}

func part1(joltages []int) int {
	distribution := joltageDistribution(joltages)
	return distribution[1] * distribution[3]
}

func compatibleJoltage(a int, b int) bool {
	return -3 <= b-a && b-a <= 3
}

func prependInt(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}

func part2(inputJoltage int, joltages []int, cache map[int]int) int {
	joltages = append(joltages, builtInJoltage(joltages))
	joltages = prependInt(joltages, 0)
	sort.Ints(joltages)
	cache[joltages[len(joltages)-1]] = 1
	for i := len(joltages) - 1; i >= 0; i-- {
		sum := 0
		for j := i + 1; j < len(joltages); j++ {
			joltCurrent := joltages[i]
			joltPrevious := joltages[j]
			if joltPrevious-joltCurrent > 3 {
				break
			}
			if compatibleJoltage(joltCurrent, joltPrevious) {
				sum += cache[joltPrevious]
			}
			cache[joltCurrent] = sum
		}
	}
	return cache[0]
}
