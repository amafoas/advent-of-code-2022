package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	rounds := strings.Split(strings.TrimSpace(string(dat)), "\n")

	t1 := time.Now()
	scores_first := map[string]int{
		"A X": 1 + 3, "A Y": 2 + 6, "A Z": 3 + 0,
		"B X": 1 + 0, "B Y": 2 + 3, "B Z": 3 + 6,
		"C X": 1 + 6, "C Y": 2 + 0, "C Z": 3 + 3,
	}

	fmt.Printf("First part: %d\n", getScore(rounds, scores_first))
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	t2 := time.Now()
	scores_second := map[string]int{
		"A X": 3 + 0, "A Y": 1 + 3, "A Z": 2 + 6,
		"B X": 1 + 0, "B Y": 2 + 3, "B Z": 3 + 6,
		"C X": 2 + 0, "C Y": 3 + 3, "C Z": 1 + 6,
	}

	fmt.Printf("Second part: %d\n", getScore(rounds, scores_second))
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
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
