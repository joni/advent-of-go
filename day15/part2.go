package day15

func Part2(numbers []int) int {
	return solve(numbers, 30_000_000)
}

func solve(numbers []int, N int) int {
	lastSeen := make(map[int]int)
	for i, n := range numbers[:len(numbers)-1] {
		lastSeen[n] = i
	}
	prevN := numbers[len(numbers)-1]
	for i := len(numbers); i < N; i++ {
		prevIdx, ok := lastSeen[prevN]
		lastSeen[prevN] = i - 1
		if ok {
			prevN = (i - 1) - prevIdx
		} else {
			prevN = 0
		}
	}
	return prevN
}
