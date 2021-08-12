package day13

import "testing"

func TestPart2Example1(t *testing.T) {
	actual := solve2("7,13,x,x,59,x,31,19")
	expected := int64(1068781)
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2Example2(t *testing.T) {
	actual := solve2("17,x,13,19")
	expected := int64(3417)
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2Example3(t *testing.T) {
	actual := solve2("67,7,59,61")
	var expected int64 = 754018
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2Example4(t *testing.T) {
	actual := solve2("67,x,7,59,61")
	var expected int64 = 779210
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2Example5(t *testing.T) {
	actual := solve2("67,7,x,59,61")
	var expected int64 = 1261476
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2Example6(t *testing.T) {
	actual := solve2("1789,37,47,1889")
	var expected int64 = 1202161486
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	var expected int64 = 741745043105674
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
