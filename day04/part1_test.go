package day04

import (
	"testing"
)

func TestReadPassport(t *testing.T) {
	actual := readPassport(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`)
	expect := Passport{ecl: "gry", pid: "860033327", eyr: 2020, hcl: "#fffffd",
		byr: 1937, iyr: 2017, cid: "147", hgt: "183cm"}
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}

func TestIsValid1(t *testing.T) {
	actual := readPassport(`hcl:#cfa07d eyr:2025 pid:166559648
	iyr:2011 ecl:brn hgt:59in`).isValid()
	expect := false
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}

func TestPart1Example(t *testing.T) {
	actual := CountValid(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`)
	expect := 2
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expect := 230
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}
