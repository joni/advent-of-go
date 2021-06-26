package day04

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func (pport Passport) isStrictValid() bool {
	return pport.isValid() &&
		isNumberInRange(pport.byr, 1920, 2002) &&
		isNumberInRange(pport.iyr, 2010, 2020) &&
		isNumberInRange(pport.eyr, 2020, 2030) &&
		isValidHgt(pport.hgt) &&
		isValidHcl(pport.hcl) &&
		isValidEcl(pport.ecl) &&
		isValidPid(pport.pid)

}

func isNumberInRange(value int, min int, max int) bool {
	return min <= value && value <= max
}

func isValidHgt(hgt string) bool {
	hgtRe := regexp.MustCompile(`([0-9]+)(cm|in)`)
	parts := hgtRe.FindStringSubmatch(hgt)
	if len(parts) != 3 {
		return false
	}
	number, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	switch parts[2] {
	case "cm":
		return isNumberInRange(number, 150, 193)
	case "in":
		return isNumberInRange(number, 59, 76)
	default:
		return false
	}
}

func isValidHcl(hcl string) bool {
	match, err := regexp.MatchString(`#[0-9a-f]{6}`, hcl)
	if err != nil {
		panic(err)
	}
	return match
}

func isValidEcl(ecl string) bool {
	match, err := regexp.MatchString(`amb|blu|brn|gry|grn|hzl|oth`, ecl)
	if err != nil {
		panic(err)
	}
	return match
}

func isValidPid(pid string) bool {
	match, err := regexp.MatchString(`^[0-9]{9}$`, pid)
	if err != nil {
		panic(err)
	}
	return match
}

func CountStrictValid(passports string) int {
	blocks := strings.Split(passports, "\n\n")
	countValid := 0
	for _, block := range blocks {
		pport := readPassport(block)
		if pport.isStrictValid() {
			countValid++
		}
	}
	return countValid
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return CountStrictValid(string(content))
}
