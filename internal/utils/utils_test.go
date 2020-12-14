package utils

import (
	"testing"
)

func TestReadLines(t *testing.T) {
	expected := []string{"a", "b", "c", ""}
	parsedInput := ReadLines("test_file_1")
	for i := range expected {
		if expected[i] != parsedInput[i] {
			t.Errorf("input parse fail at index %d", i)
		}
	}
}

func TestReadLinesNoBlankLine(t *testing.T) {
	expected := []string{"a", "b", "c"}
	parsedInput := ReadLines("test_file_2")
	for i := range expected {
		if expected[i] != parsedInput[i] {
			t.Errorf("input parse fail at index %d", i)
		}
	}
}

func TestAdjacentCells(t *testing.T) {
	var tests = []struct {
		row      int
		col      int
		maxRow   int
		maxCol   int
		expected [][]int
	}{
		{0, 0, 0, 0, [][]int{}},
		{0, 0, 1, 1, [][]int{{0, 1}, {1, 0}, {1, 1}}},
		{0, 0, 1, 0, [][]int{{1, 0}}},
		{0, 0, 0, 1, [][]int{{0, 1}}},
		{3, 3, 2, 2, [][]int{{2, 2}}},
		{0, 0, 2, 2, [][]int{{0, 1}, {1, 0}, {1, 1}}},
		{0, 1, 2, 2, [][]int{{0, 0}, {0, 2}, {1, 0}, {1, 1}, {1, 2}}},
		{1, 1, 2, 2, [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
		{2, 2, 2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}}},
		{0, 2, 9, 9, [][]int{{0, 1}, {0, 3}, {1, 1}, {1, 2}, {1, 3}}},
	}

	for _, test_case := range tests {
		t.Run("", func(t *testing.T) {
			result := AdjacentCells(test_case.row, test_case.col, test_case.maxRow, test_case.maxCol)
			if len(result) != len(test_case.expected) {
				t.Errorf("invalid length, got: %d, expected %d", len(result), len(test_case.expected))
			}
			for i := range result {
				if len(result) != len(test_case.expected) {
					t.Errorf("invalid length at index %d, got: %d, expected %d", i, len(result), len(test_case.expected))
				}
				for j := range result[i] {
					if result[i][j] != test_case.expected[i][j] {
						t.Errorf("invalid result %d, %d", result[i][j], test_case.expected[i][j])
					}
				}
			}
		})
	}
}
