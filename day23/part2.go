package day23

import (
	"container/list"
	"fmt"
)

func solve2(labels []int) int {
	cups := list.New()
	elemMap := make(map[int]*list.Element)
	maxLabel := 1000_000
	for _, label := range labels {
		elemMap[label] = cups.PushBack(label)
	}
	for label := 10; label <= maxLabel; label++ {
		elemMap[label] = cups.PushBack(label)
	}

	moves := 10_000_000
	for i := 0; i < moves; i++ {
		step(cups, elemMap, maxLabel)
	}

	for cups.Front().Value != 1 {
		cups.MoveToBack(cups.Front())
	}
	cups.MoveToBack(cups.Front())

	label1 := cups.Front().Value.(int)
	label2 := cups.Front().Next().Value.(int)

	fmt.Println("labels after 1 are", label1, label2)
	return label1 * label2
}
