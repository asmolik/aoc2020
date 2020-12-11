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
