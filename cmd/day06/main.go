package main

import (
	"fmt"

	"github.com/asmolik/aoc2020/internal/utils"
)

func parseInput(lines []string) [][]string {
	ret := [][]string{}
	ret = append(ret, []string{})
	current_group := 0
	for _, line := range lines {
		if line == "" {
			ret = append(ret, []string{})
			current_group++
			continue
		}
		ret[current_group] = append(ret[current_group], line)
	}
	return ret
}

func main() {
	data_str := utils.ReadLines("input")
	answers := parseInput(data_str)
	fmt.Println(part1(answers))
	fmt.Println(part2(answers))
}

func part1(answers [][]string) int {
	sum := 0
	for _, group := range answers {
		a := make(map[rune]bool)
		for _, person := range group {
			for _, question := range person {
				a[question] = true
			}
		}
		sum += len(a)
	}
	return sum
}

func part2(answers [][]string) int {
	sum := 0
	for _, group := range answers {
		a := make(map[rune]int)
		for _, person := range group {
			for _, question := range person {
				a[question] = a[question] + 1
			}
		}
		for _, v := range a {
			if v == len(group) {
				sum++
			}
		}
	}
	return sum
}
