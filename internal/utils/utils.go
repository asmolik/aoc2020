package utils

import (
	"io/ioutil"
	"strings"
)

func ReadLines(filename string) []string {
	test_data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(test_data), "\n")
}

func AdjacentCells(row int, col int, maxRow int, maxCol int) [][2]int {
	ret := [][2]int{}
	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i > maxRow {
			continue
		}
		for j := col - 1; j <= col+1; j++ {
			if j < 0 || j > maxCol {
				continue
			}
			if i == row && j == col {
				continue
			}
			ret = append(ret, [2]int{i, j})
		}
	}
	return ret
}
