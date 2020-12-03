package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	min, max  int
	character rune
	password  []rune
}

func readInput(lines []string) []Password {
	passwords := []Password{}
	r, _ := regexp.Compile("([\\d]*)-([\\d]*) (.): (.*)")
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		min, _ := strconv.Atoi(match[1])
		max, _ := strconv.Atoi(match[2])
		passwords = append(passwords, Password{min, max, []rune(match[3])[0], []rune(match[4])})
	}
	return passwords
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	data_str := strings.Split(string(data), "\n")
	passwords := readInput(data_str)
	fmt.Println(part1(passwords))
	fmt.Println(part2(passwords))
}

func part1(passwords []Password) int {
	error_count := 0
	for _, pass := range passwords {
		occurences := 0
		for _, i := range pass.password {
			if i == pass.character {
				occurences++
			}
		}
		if occurences > pass.max || occurences < pass.min {
			error_count++
		}
	}
	return len(passwords) - error_count
}

func part2(passwords []Password) int {
	error_count := 0
	for _, pass := range passwords {
		if pass.password[pass.min-1] == pass.character && pass.password[pass.max-1] == pass.character {
			error_count++
        }
        if pass.password[pass.min-1] == pass.character || pass.password[pass.max-1] == pass.character {
            continue
        }
        error_count++
	}
	return len(passwords) - error_count
}
