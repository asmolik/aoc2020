package main

import (
	"testing"

	"github.com/asmolik/aoc2020/internal/utils"
)

func TestPart1(t *testing.T) {
	in1, in2 := parseInput(utils.ReadLines("test_input"))
	var tests = []struct {
		name     string
		input1   int
		input2   []int
		expected int
	}{
		{"input_1", in1, in2, 295},
	}

	for _, test_case := range tests {
		testname := test_case.name
		t.Run(testname, func(t *testing.T) {
			result := part1(test_case.input1, test_case.input2)
			if result != test_case.expected {
				t.Errorf("got %d, want %d", result, test_case.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	_, in2 := parseInput(utils.ReadLines("test_input"))
	var tests = []struct {
		name     string
		input    []int
		expected int
	}{
		{"input_1", in2, 1068781},
		{"input_1", []int{17, 0, 13, 19}, 3417},
		{"input_1", []int{67, 7, 59, 61}, 754018},
		{"input_1", []int{67, 0, 7, 59, 61}, 779210},
		{"input_1", []int{67, 7, 0, 59, 61}, 1261476},
		{"input_1", []int{1789, 37, 47, 1889}, 1202161486},
	}

	for _, test_case := range tests {
		testname := test_case.name
		t.Run(testname, func(t *testing.T) {
			result := part2(test_case.input)
			if result != test_case.expected {
				t.Errorf("got %d, want %d", result, test_case.expected)
			}
		})
	}
}
