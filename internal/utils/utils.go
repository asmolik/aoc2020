package utils

import (
	"io/ioutil"
	"strings"
)

func ReadLines(filename string) []string {
	test_data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(test_data), "\n")
}
