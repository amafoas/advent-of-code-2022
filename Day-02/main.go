package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	rounds := strings.Split(strings.TrimSpace(string(dat)), "\n")

	scores_first := map[string]int{
		"A X": 1 + 3, "A Y": 2 + 6, "A Z": 3 + 0,
		"B X": 1 + 0, "B Y": 2 + 3, "B Z": 3 + 6,
		"C X": 1 + 6, "C Y": 2 + 0, "C Z": 3 + 3,
	}

	fmt.Println("First part: ", getScore(rounds, scores_first))

	scores_second := map[string]int{
		"A X": 3 + 0, "A Y": 1 + 3, "A Z": 2 + 6,
		"B X": 1 + 0, "B Y": 2 + 3, "B Z": 3 + 6,
		"C X": 2 + 0, "C Y": 3 + 3, "C Z": 1 + 6,
	}

	fmt.Println("Second part: ", getScore(rounds, scores_second))
}

func getScore(rounds []string, scores map[string]int) int {
	totalScore := 0
	for _, round := range rounds {
		score, ok := scores[round]
		if ok {
			totalScore += score
		}
	}
	return totalScore
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
