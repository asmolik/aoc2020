package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/asmolik/aoc2020/internal/utils"
)

type PasswordVerifier interface {
	verify(string) bool
}

type RangeVerifier struct {
	min    string
	max    string
	length int
}

type Verification func(PasswordVerifier, string, bool) bool

func (verifier RangeVerifier) verify(value string) bool {
	if len(value) != verifier.length {
		return false
	}
	return verifier.min <= value && value <= verifier.max
}

type HeightVerifier struct{}

func (verifier HeightVerifier) verify(value string) bool {
	if strings.HasSuffix(value, "in") {
		return "59in" <= value && value <= "76in"
	} else if strings.HasSuffix(value, "cm") {
		return "150cm" <= value && value <= "193cm"
	}
	return false
}

type RegexVerifier struct {
	regex regexp.Regexp
}

func (verifier RegexVerifier) verify(value string) bool {
	return verifier.regex.MatchString(value)
}

func parsePassports(lines []string, verification Verification) int {
	validPassports := 0
	passportIndex := 0
	passportLineIndex := 0
	// currPassportLineIndex := 0
	currPassport := make(map[string]string)
	for _, line := range lines {
		passportLineIndex++
		if line == "" {
			if verifyPassport(currPassport, verification) {
				validPassports++
			} //else {
			// 	fmt.Println("passport invalid " + strconv.Itoa(currPassportLineIndex) + " " + strconv.Itoa(passportIndex))
			// }
			currPassport = make(map[string]string)
			passportIndex++
			continue
		}
		// if len(currPassport) == 0 {
		// 	currPassportLineIndex = passportLineIndex
		// }
		for k, v := range parseLine(line) {
			currPassport[k] = v
		}
	}
	return validPassports
}

func parseLine(line string) map[string]string {
	passport := make(map[string]string)
	fields := strings.Fields(string(line))
	for _, field := range fields {
		tmp := strings.Split(field, ":")
		passport[tmp[0]] = tmp[1]
	}
	return passport
}

func verifyPassport(passport map[string]string, verification Verification) bool {
	hcl_regex, _ := regexp.Compile("^#([0-9a-f]){6}$")
	ecl_regex, _ := regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	pid_regex, _ := regexp.Compile("^\\d{9}$")
	requiredFields := map[string]PasswordVerifier{
		"byr": RangeVerifier{"1920", "2002", 4},
		"iyr": RangeVerifier{"2010", "2020", 4},
		"eyr": RangeVerifier{"2020", "2030", 4},
		"hgt": HeightVerifier{},
		"hcl": RegexVerifier{*hcl_regex},
		"ecl": RegexVerifier{*ecl_regex},
		"pid": RegexVerifier{*pid_regex},
	}
	for key, field := range requiredFields {
		value, present := passport[key]
		if !verification(field, value, present) {
			// fmt.Println("field invalid " + key + " " + strconv.FormatBool(present) + " " + value)
			return false
		}
	}
	return true
}

func main() {
	test_data_str := utils.ReadLines("test_input")
	test_data_valid_passports_str := utils.ReadLines("valid_passports")
	test_data_invalid_passports_str := utils.ReadLines("invalid_passports")
	fmt.Println(parsePassports(test_data_str, part1))
	fmt.Println(parsePassports(test_data_valid_passports_str, part2))
	fmt.Println(parsePassports(test_data_invalid_passports_str, part2))
	data_str := utils.ReadLines("input")
	fmt.Println(parsePassports(data_str, part1))
	fmt.Println(parsePassports(data_str, part2))
}

func part1(verifier PasswordVerifier, value string, present bool) bool {
	return present
}

func part2(verifier PasswordVerifier, value string, present bool) bool {
	return present && verifier.verify(value)
}
