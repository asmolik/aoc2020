package main

import (
	"testing"

	"github.com/asmolik/aoc2020/internal/utils"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		name     string
		input    Ferry
		expected int
	}{
		{"input_1", parseInput(utils.ReadLines("test_input")), 37},
	}

	for _, test_case := range tests {
		testname := test_case.name
		t.Run(testname, func(t *testing.T) {
			result := part1(test_case.input)
			if result != test_case.expected {
				t.Errorf("got %d, want %d", result, test_case.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		name     string
		input    Ferry
		expected int
	}{
		{"input_1", parseInput(utils.ReadLines("test_input")), 26},
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
