package main

import (
	"fmt"
	"math"
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
	data_str := utils.ReadLines("test_input")
	run_for_data(data_str, 5)
	data_str = utils.ReadLines("input")
	run_for_data(data_str, 25)
}

func run_for_data(data []string, preamble int) {
	numbers := parseInput(data)
	xmas := part1(numbers, preamble)
	fmt.Println(xmas)
	fmt.Println(part2(numbers, preamble, xmas))
}

func part1(numbers []int, preamble int) int {
	candidates := map[int]bool{}
	for _, n := range numbers[:preamble] {
		candidates[n] = true
	}
	for idx, n := range numbers[preamble:] {
		found := false
		for i := preamble; i > 0; i-- {
			tmp := numbers[idx+preamble-i]
			if _, exists := candidates[n-tmp]; exists && 2*tmp != n {
				delete(candidates, numbers[idx])
				candidates[n] = true
				found = true
				break
			}
		}
		if !found {
			return n
		}
	}
	return -1
}

func part2(numbers []int, preamble int, xmas int) int {
	sumStart := 0
	sumEnd := 0
	for idx, n := range numbers {
		xmas -= n
		for xmas < 0 {
			xmas += numbers[sumStart]
			sumStart++
			if xmas == 0 {
				sumEnd = idx
				break
			}
		}
		if xmas == 0 {
			sumEnd = idx
			break
		}
	}
	min := math.MaxInt64
	max := -1
	for i := sumStart; i <= sumEnd; i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return min + max
}
