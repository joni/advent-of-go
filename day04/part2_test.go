package day04

import (
	"testing"
)

func TestNumberInRange(t *testing.T) {
	actual := isNumberInRange(2020, 2010, 2020) &&
		isNumberInRange(2010, 2010, 2020) &&
		!isNumberInRange(2021, 2010, 2020) &&
		!isNumberInRange(2009, 2010, 2020)
	expect := true
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}

func TestValidHgt(t *testing.T) {
	actual := isValidHgt("60in") &&
		isValidHgt("190cm") &&
		!isValidHgt("190in") &&
		!isValidHgt("190")
	expect := true
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}

func TestValidHcl(t *testing.T) {
	actual := isValidHcl("#123abc") &&
		!isValidHcl("#123abz") &&
		!isValidHcl("123abc")
	expect := true
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}

func TestValidEcl(t *testing.T) {
	actual := isValidEcl("brn") &&
		!isValidEcl("wat")
	expect := true
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}

func TestValidPid(t *testing.T) {
	actual := isValidPid("000000001") &&
		!isValidPid("0123456789")
	expect := true
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}

func TestStrictValid(t *testing.T) {
	actual := CountStrictValid(`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
	`)
	expect := 4
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}
func TestStrictInvalid(t *testing.T) {
	actual := CountStrictValid(`eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`)
	expect := 0
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	expect := 156
	if actual != expect {
		t.Errorf("actual %v expected %v", actual, expect)
	}
}
