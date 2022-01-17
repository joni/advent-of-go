package day22

import (
	"io/ioutil"
	"strings"
)

type GameSet map[int]struct{}

func NewGameSet() GameSet {
	return make(GameSet)
}

func (games GameSet) Add(cards1, cards2 []int) bool {
	hash := 0
	for _, c := range cards1 {
		hash = hash*97 + c
	}
	for _, c := range cards2 {
		hash = hash*91 + c
	}
	if _, ok := games[hash]; ok {
		return false
	} else {
		games[hash] = struct{}{}
		return true
	}
}

func winner(player1, player2 []int) int {
	if player1[0] < len(player1) && player2[0] < len(player2) {
		// recursive
		subPlayer1 := make([]int, player1[0])
		subPlayer2 := make([]int, player2[0])
		copy(subPlayer1, player1[1:1+player1[0]])
		copy(subPlayer2, player2[1:1+player2[0]])
		winner, _ := play2(subPlayer1, subPlayer2)
		return winner
	} else if player1[0] > player2[0] {
		return 1
	} else {
		return 2
	}
}

func play2(player1, player2 []int) (int, int) {
	games := NewGameSet()
	for len(player1) > 0 && len(player2) > 0 {
		if !games.Add(player1, player2) {
			// infinite game! player 1 wins
			break
		}
		switch winner(player1, player2) {
		case 1:
			player1 = append(player1, player1[0], player2[0])
		case 2:
			player2 = append(player2, player2[0], player1[0])
		}
		player1 = player1[1:]
		player2 = player2[1:]
	}

	if len(player1) > 0 {
		return 1, score(player1)
	} else {
		return 2, score(player2)
	}
}

func solve2(input string) int {
	players := strings.Split(input, "\n\n")
	player1 := parsePlayer(players[0])
	player2 := parsePlayer(players[1])
	_, score := play2(player1, player2)
	return score
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content))
}
