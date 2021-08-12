package day13

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Bus struct {
	offset    int
	frequency int
}

func solve2(input string) int64 {
	inputSplit := strings.Split(input, ",")

	buses := []Bus{}
	N := int64(1)
	for idx, busStr := range inputSplit {
		if busStr != "x" {
			busId, _ := strconv.Atoi(busStr)
			buses = append(buses, Bus{idx, busId})
			N *= int64(busId)
		}
	}
	fmt.Println(buses, N)

	// now find number T such that T + offset is divisible by frequency for each bus
	// T = -o[i] (mod f[i]). The solution takes the form T = sum(e[i]*o[i]) where
	// e[i] = -1 (mod f[i]) and e[i] = 0 (mod f[j]) for j != i
	// Set d[i] = prod(f)/f[i]: e[i] = -d[i]*modinv(d[i], f[i])

	T := int64(0)
	for _, bus := range buses {
		d := N / int64(bus.frequency)
		e := -d * modinv(d, int64(bus.frequency))
		T += int64(bus.offset) * e % N
	}
	return (N + T%N) % N
}

func modinv(d, f int64) int64 {
	for c := int64(1); c < f; c++ {
		if c*d%f == 1 {
			return c
		}
	}
	return 0
}

func Part2() int64 {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	return solve2(lines[1])

}
