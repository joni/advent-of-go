package day21

import "testing"

const example = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func Test_Part1Example(t *testing.T) {
	actual := solve1(example)
	expected := 5
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part1(t *testing.T) {
	actual := Part1()
	expected := 2584
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
