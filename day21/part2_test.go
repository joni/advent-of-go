package day21

import "testing"

func Test_Part2Example(t *testing.T) {
	actual := solve2(example)
	expected := "mxmxvkd,sqjhc,fvjkl"
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2()
	expected := "fqhpsl,zxncg,clzpsl,zbbnj,jkgbvlxh,dzqc,ppj,glzb"
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
