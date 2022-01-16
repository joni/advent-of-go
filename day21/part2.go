package day21

import (
	"io/ioutil"
	"sort"
	"strings"
)

func solve2(input string) string {
	matrix, _ := analyze(input)

	var allergens []string
	for a := range matrix {
		allergens = append(allergens, a)
	}

	sort.Strings(allergens)

	var indredientsWithAllergens []string
	for _, a := range allergens {
		for i := range matrix[a] {
			indredientsWithAllergens = append(indredientsWithAllergens, i)
		}
	}

	return strings.Join(indredientsWithAllergens, ",")
}

func Part2() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(strings.TrimSpace(string(content)))
}
