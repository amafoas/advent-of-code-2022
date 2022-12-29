package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	elfs := make(map[Pos]bool)
	for y, line := range lines {
		row := strings.Split(line, "")
		for x, c := range row {
			if c == "#" {
				p := Pos{x, len(lines) - y - 1}
				elfs[p] = true
			}
		}
	}
	t1 := time.Now()
	p1, _ := simulate(10, elfs)
	fmt.Printf("First part: %d\n", getArea(p1))
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	t2 := time.Now()
	_, p2 := simulate(2000, elfs)
	fmt.Printf("First part: %d\n", p2)
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
}

func simulate(t int, elfs map[Pos]bool) (map[Pos]bool, int) {
	old := elfs
	var stable int
	for i := 0; i < t; i++ {
		moves := make(map[Pos][]Pos)
		for e := range old {
			np := getDirection(e, i, old)
			moves[np] = append(moves[np], e)
		}
		new := make(map[Pos]bool)
		for p, arr := range moves {
			if len(arr) == 1 {
				new[p] = true
			} else {
				for _, e := range arr {
					new[e] = true
				}
			}
		}
		if mapsEquals(old, new) {
			stable = i + 1
			break
		}
		old = new
	}
	return old, stable
}

func mapsEquals(m, r map[Pos]bool) bool {
	res := len(m) == len(r)
	if res {
		for k, v := range m {
			if r[k] != v {
				res = false
				break
			}
		}
	}
	return res
}

func getArea(m map[Pos]bool) int {
	var maxX, minX int = 0, math.MaxInt
	var maxY, minY int = 0, math.MaxInt
	for k := range m {
		maxX = max(maxX, k.x)
		maxY = max(maxY, k.y)
		minX = min(minX, k.x)
		minY = min(minY, k.y)
	}

	return ((maxX - minX + 1) * (maxY - minY + 1)) - len(m)
}

func getDirection(e Pos, off int, mp map[Pos]bool) Pos {
	type Dir struct {
		can bool
		pos Pos
	}
	moves := []Dir{
		{!mp[Pos{e.x, e.y + 1}] && !mp[Pos{e.x + 1, e.y + 1}] && !mp[Pos{e.x - 1, e.y + 1}], Pos{e.x, e.y + 1}}, // N
		{!mp[Pos{e.x, e.y - 1}] && !mp[Pos{e.x + 1, e.y - 1}] && !mp[Pos{e.x - 1, e.y - 1}], Pos{e.x, e.y - 1}}, // S
		{!mp[Pos{e.x - 1, e.y}] && !mp[Pos{e.x - 1, e.y + 1}] && !mp[Pos{e.x - 1, e.y - 1}], Pos{e.x - 1, e.y}}, // W
		{!mp[Pos{e.x + 1, e.y}] && !mp[Pos{e.x + 1, e.y + 1}] && !mp[Pos{e.x + 1, e.y - 1}], Pos{e.x + 1, e.y}}, // E
	}

	res := e
	if !(moves[0].can && moves[1].can && moves[2].can && moves[3].can) {
		for i := off; i < off+4; i++ {
			if m := moves[i%4]; m.can {
				res = m.pos
				break
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
