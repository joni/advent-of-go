package day21

import (
	"bufio"
	"io/ioutil"
	"regexp"
	"strings"
)

type WordCounter map[string]int

func NewCounter() WordCounter {
	return make(WordCounter)
}

func (counter WordCounter) Add(word string) bool {
	if _, exists := counter[word]; exists {
		counter[word]++
		return false
	} else {
		counter[word] = 1
		return true
	}
}

type WordSet map[string]int

func NewSet() WordSet {
	return make(WordSet)
}

func (set WordSet) Add(word string) bool {
	if _, exists := set[word]; !exists {
		set[word] = len(set)
		return true
	} else {
		return false
	}
}

func (set WordSet) Copy() WordSet {
	newSet := NewSet()
	for key, value := range set {
		newSet[key] = value
	}
	return newSet
}

func (set WordSet) RetainAll(otherSet WordSet) {
	for key := range set {
		if _, ok := otherSet[key]; !ok {
			delete(set, key)
		}
	}
}

func (set WordSet) RemoveAll(otherSet WordSet) bool {
	deleted := false
	for key := range otherSet {
		if _, ok := set[key]; ok {
			delete(set, key)
			deleted = true
		}
	}
	return deleted
}

func removeFromValues(matrix map[string]WordSet, excluded string, words WordSet) bool {
	modified := false
	for key, value := range matrix {
		if key != excluded {
			if value.RemoveAll(words) {
				modified = true
			}
		}
	}
	return modified
}

func analyze(input string) (map[string]WordSet, WordCounter) {
	wordRe := regexp.MustCompile(`\w+`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	uniqIngredients := NewCounter()
	matrix := make(map[string]WordSet)
	for scanner.Scan() {
		line := scanner.Text()
		match := wordRe.FindAllString(line, -1)
		containsRead := false
		ingredients := NewSet()
		for _, word := range match {
			if word == "contains" {
				containsRead = true
				continue
			}
			if containsRead {
				if _, ok := matrix[word]; ok {
					matrix[word].RetainAll(ingredients)
				} else {
					matrix[word] = ingredients.Copy()
				}
			} else {
				uniqIngredients.Add(word)
				ingredients.Add(word)
			}
		}
	}

	loop := true
	for loop {
		loop = false
		for allergen, ingredients := range matrix {
			if len(ingredients) == 1 {
				if removeFromValues(matrix, allergen, ingredients) {
					loop = true
				}
			}
		}
	}

	return matrix, uniqIngredients
}

func solve1(input string) int {
	matrix, uniqIngredients := analyze(input)

	indredientsWithAllergens := NewSet()
	for _, ingredients := range matrix {
		for i := range ingredients {
			indredientsWithAllergens.Add(i)
		}
	}

	count := 0
	for ingr, occurrances := range uniqIngredients {
		if _, ok := indredientsWithAllergens[ingr]; !ok {
			count += occurrances
		}
	}
	return count
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve1(strings.TrimSpace(string(content)))
}
