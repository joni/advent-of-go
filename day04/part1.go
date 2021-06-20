package day04

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func readPassport(str string) Passport {
	var pport Passport
	fields := strings.Fields(str)
	for _, entry := range fields {
		keyval := strings.SplitN(entry, ":", 2)
		key, val := keyval[0], keyval[1]
		switch key {
		case "byr":
			pport.byr, _ = strconv.Atoi(val)
		case "iyr":
			pport.iyr, _ = strconv.Atoi(val)
		case "eyr":
			pport.eyr, _ = strconv.Atoi(val)
		case "hgt":
			pport.hgt = val
		case "hcl":
			pport.hcl = val
		case "ecl":
			pport.ecl = val
		case "pid":
			pport.pid = val
		case "cid":
			pport.cid = val

		}
	}
	return pport
}

func (pport Passport) isValid() bool {
	return pport.byr != 0 && pport.iyr != 0 && pport.eyr != 0 &&
		pport.hgt != "" && pport.hcl != "" &&
		pport.ecl != "" && pport.pid != ""
}

func CountValid(passports string) int {
	blocks := strings.Split(passports, "\n\n")
	countValid := 0
	for _, block := range blocks {
		if readPassport(block).isValid() {
			countValid++
		}
	}
	return countValid
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return CountValid(string(content))
}
