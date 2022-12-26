package main

import (
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	x, y int
}

func addPair(a, b Pair) Pair {
	return Pair{a.x + b.x, a.y + b.y}
}

type Rock struct {
	size Pair
	body []Pair
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	ins := strings.Split(strings.TrimSpace(string(dat)), "")

	fmt.Println("First part: ", simulate(2022, ins))
}

func simulate(n int, ins []string) int {
	rocks := []Rock{
		{Pair{4, 1}, []Pair{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},         // -
		{Pair{3, 3}, []Pair{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}}, // +
		{Pair{3, 3}, []Pair{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}}, // L
		{Pair{1, 4}, []Pair{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},         // |
		{Pair{2, 2}, []Pair{{0, 0}, {1, 0}, {0, 1}, {1, 1}}},         // #
	}

	chamber := make(map[Pair]bool)
	var top int // top of rocks

	i := -1
	for r := 0; r < n; r++ {
		rock := rocks[r%len(rocks)]
		pos := Pair{2, top + 3}

		for {
			i++
			mv := ins[i%len(ins)]
			if mv == ">" {
				np := addPair(pos, Pair{1, 0})
				if !hitARock(rock, np, chamber) {
					pos = np
				}
			} else {
				np := addPair(pos, Pair{-1, 0})
				if !hitARock(rock, np, chamber) {
					pos = np
				}
			}
			np := addPair(pos, Pair{0, -1})
			if !hitARock(rock, np, chamber) {
				pos = np
			} else {
				for _, p := range rock.body {
					chamber[addPair(p, pos)] = true
					top = max(top, pos.y+rock.size.y)
				}
				break
			}
		}

	}
	return top
}

func hitARock(r Rock, p Pair, m map[Pair]bool) bool {
	for _, b := range r.body {
		ap := addPair(p, b)
		if m[ap] || ap.y < 0 || ap.x < 0 || ap.x >= 7 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
