package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/asmolik/aoc2020/internal/utils"
)

type Rule struct {
	bag        string
	inner_bags map[string]int
	outer_bags map[string]bool
	visited    bool
}

func findRule(key string, rules map[string]Rule) *Rule {
	if val, ok := rules[key]; ok {
		return &val
	} else {
		rule := Rule{key, map[string]int{}, map[string]bool{}, false}
		rules[key] = rule
		return &rule
	}
}

func parseInput(lines []string) map[string]Rule {
	ret := map[string]Rule{}
	regex := regexp.MustCompile("[ ]?(\\d*) (.*) bag")
	for _, line := range lines {
		a := strings.Split(line, " bags contain ")
		outer_bag := a[0]
		if a[1] == "no other bags." {
			findRule(outer_bag, ret)
			continue
		}
		other_bags := strings.Split(a[1], ",")
		inner_bags := map[string]int{}
		for _, bag := range other_bags {
			match := regex.FindStringSubmatch(bag)
			count, _ := strconv.Atoi(match[1])
			color := match[2]
			inner_bags[color] = count
			rule := findRule(color, ret)
			rule.outer_bags[outer_bag] = true
		}
		rule := findRule(outer_bag, ret)
		for k, v := range inner_bags {
			rule.inner_bags[k] = v
		}
	}
	return ret
}

func countOuterColors(color string, rules map[string]Rule, visited map[string]bool) int {
	sum := 0
	rule := rules[color]
	if rule.visited {
		return 0
	}
	rule.visited = true
	for col := range rule.outer_bags {
		if _, ok := visited[col]; ok {
			continue
		}
		visited[col] = true
		sum++
		sum += countOuterColors(col, rules, visited)
	}
	return sum
}

func countInnerBags(color string, rules map[string]Rule) int {
	sum := 0
	rule := rules[color]
	for col, count := range rule.inner_bags {
		// fmt.Println(color + " contains " + strconv.Itoa(count) + " " + col)
		sum += count
		sum += count * countInnerBags(col, rules)
	}
	return sum
}

func main() {
	data_str := utils.ReadLines("test_input")
	run_for_data(data_str)
	data_str = utils.ReadLines("input")
	run_for_data(data_str)
}

func run_for_data(data []string) {
	rules := parseInput(data)
	fmt.Println(part1(rules))
	fmt.Println(part2(rules))

}

func part1(rules map[string]Rule) int {
	const shiny_gold = "shiny gold"
	return countOuterColors(shiny_gold, rules, map[string]bool{})
}

func part2(rules map[string]Rule) int {
	const shiny_gold = "shiny gold"
	return countInnerBags(shiny_gold, rules)
}
