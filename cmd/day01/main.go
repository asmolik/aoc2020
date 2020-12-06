package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	data_str := strings.Split(string(data), "\n")
	numbers := []int{}
	for _, i := range data_str {
        if i != "" {
            a, _ := strconv.Atoi(i)
            numbers = append(numbers, a)
        }
    }
    fmt.Println(part1(numbers))
    fmt.Println(part2(numbers))
}

func part1(numbers []int) int {
    map2020 := make(map[int]bool)
	for _, i := range numbers {
        map2020[i] = true
        if map2020[2020-i] {
            return i*(2020-i)
        }
    }
    return -1
}

func part2(numbers []int) int {
    map2020 := make(map[int]bool)
	for _, i := range numbers {
        for _, j := range numbers {
            map2020[j] = true
            if map2020[2020-i-j] {
                return i*j*(2020-i-j)
            }
        }
    }
    return -1
}
