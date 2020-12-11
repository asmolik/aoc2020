package main

import (
	"testing"

	"github.com/asmolik/aoc2020/internal/utils"
)

func TestParseInput(t *testing.T) {
	data_str := utils.ReadLines("test_input_1")
	expected := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	parsedInput := parseInput(data_str)
	for i := range expected {
		if expected[i] != parsedInput[i] {
			t.Errorf("input parse fail at index %d", i)
		}
	}
}

func TestBuiltInJoltage(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected int
	}{
		{"input_1", parseInput(utils.ReadLines("test_input_1")), 22},
		{"input_2", parseInput(utils.ReadLines("test_input_2")), 52},
	}

	for _, test_case := range tests {
		testname := test_case.name
		t.Run(testname, func(t *testing.T) {
			result := builtInJoltage(test_case.input)
			if result != test_case.expected {
				t.Errorf("got %d, want %d", result, test_case.expected)
			}
		})
	}
}

func TestJoltageDistribution(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected map[int]int
	}{
		{"input_1", parseInput(utils.ReadLines("test_input_1")), map[int]int{1: 7, 3: 5}},
		{"input_2", parseInput(utils.ReadLines("test_input_2")), map[int]int{1: 22, 3: 10}},
	}

	for _, test_case := range tests {
		testname := test_case.name
		t.Run(testname, func(t *testing.T) {
			result := joltageDistribution(test_case.input)
			for k, v := range test_case.expected {
				if v != result[k] {
					t.Errorf("got %d, want %d", result[k], v)
				}
			}
		})
	}
}

func TestPart1(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected int
	}{
		{"input_1", parseInput(utils.ReadLines("test_input_1")), 35},
		{"input_2", parseInput(utils.ReadLines("test_input_2")), 220},
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
		input    []int
		expected int
	}{
		{"input_1", parseInput(utils.ReadLines("test_input_1")), 8},
		{"input_2", parseInput(utils.ReadLines("test_input_2")), 19208},
	}

	for _, test_case := range tests {
		testname := test_case.name
		t.Run(testname, func(t *testing.T) {
			// test_case.input = append(test_case.input, builtInJoltage(test_case.input))
			result := part2(0, test_case.input, map[int]int{})
			if result != test_case.expected {
				t.Errorf("got %d, want %d", result, test_case.expected)
			}
		})
	}
}
