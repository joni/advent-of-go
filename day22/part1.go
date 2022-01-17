package day22

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parsePlayer(p string) []int {
	var cards []int
	for _, line := range strings.Split(p, "\n") {
		card, err := strconv.Atoi(line)
		if err == nil {
			cards = append(cards, card)
		}
	}
	return cards
}

func score(cards []int) int {
	score := 0
	for idx, card := range cards {
		score += (len(cards) - idx) * card
	}
	return score
}

func play(player1 []int, player2 []int) int {
	for len(player1) > 0 && len(player2) > 0 {
		// fmt.Println(player1, player2)
		if player1[0] > player2[0] {
			player1 = append(player1, player1[0], player2[0])
		} else {
			player2 = append(player2, player2[0], player1[0])
		}
		player1 = player1[1:]
		player2 = player2[1:]
	}
	fmt.Println(player1, player2)

	if len(player1) > 0 {
		return score(player1)
	} else {
		return score(player2)
	}
}

func solve(input string) int {
	players := strings.Split(input, "\n\n")
	player1 := parsePlayer(players[0])
	player2 := parsePlayer(players[1])
	return play(player1, player2)
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(string(content))
}
