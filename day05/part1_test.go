package day05

import "testing"

func TestDecodeSeat(t *testing.T) {
	actual := decodeSeat("BFFFBBFRRR")
	expected := 567
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 822
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
