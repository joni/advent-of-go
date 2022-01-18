package day23

import (
	"container/list"
	"fmt"
)

func in(num int, labels []*list.Element) bool {
	for _, label := range labels {
		if num == label.Value {
			return true
		}
	}
	return false
}

func display(lst list.List) {
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func step(cups *list.List, elemMap map[int]*list.Element, highest int) {

	// display(*cups)

	currentElem := cups.Front()
	elem1 := currentElem.Next()
	elem2 := elem1.Next()
	elem3 := elem2.Next()
	pick := []*list.Element{elem1, elem2, elem3}

	destination := currentElem.Value.(int) - 1
	if destination <= 0 {
		destination = highest
	}
	for in(destination, pick) {
		destination--
		if destination <= 0 {
			destination = highest
		}
	}

	cups.MoveAfter(pick[0], elemMap[destination])
	cups.MoveAfter(pick[1], pick[0])
	cups.MoveAfter(pick[2], pick[1])

	cups.MoveToBack(currentElem)
}

func solve(labels []int, moves int) []int {
	cups := list.New()
	elemMap := make(map[int]*list.Element)
	maxLabel := 0
	for _, label := range labels {
		elemMap[label] = cups.PushBack(label)
		if label > maxLabel {
			maxLabel = label
		}
	}

	for i := 0; i < moves; i++ {
		step(cups, elemMap, maxLabel)
	}

	for cups.Front().Value != 1 {
		cups.MoveToBack(cups.Front())
	}

	var result []int
	for e := cups.Front().Next(); e != nil; e = e.Next() {
		result = append(result, e.Value.(int))
	}

	return result
}
