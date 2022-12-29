package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pair struct {
	x int
	y int
}
type Rope struct {
	knots []Pair
	size  int
}

func (r Rope) tail() Pair {
	return r.knots[r.size-1]
}
func (r *Rope) pull(motion Pair) {
	r.knots[0].x += motion.x
	r.knots[0].y += motion.y
	for i := 1; i < r.size; i++ {
		f, s := r.knots[i-1], r.knots[i]
		dis := Pair{f.x - s.x, f.y - s.y}
		if abs(dis.x) > 1 || abs(dis.y) > 1 {
			if dis.x != 0 {
				r.knots[i].x += sign(dis.x)
			}
			if dis.y != 0 {
				r.knots[i].y += sign(dis.y)
			}
		}
	}
}

func newRope(size int) Rope {
	return Rope{make([]Pair, size), size}
}

func simulate(moves []string, ropeLen int) int {
	motions := map[string]Pair{"U": {0, 1}, "D": {0, -1}, "L": {-1, 0}, "R": {1, 0}}
	visited := make(map[Pair]bool)
	rope := newRope(ropeLen)
	for _, move := range moves {
		mv := strings.Split(move, " ")
		motion, total := motions[mv[0]], atoi(mv[1])

		for i := 0; i < total; i++ {
			rope.pull(motion)

			if !visited[rope.tail()] {
				visited[rope.tail()] = true
			}
		}
	}
	return len(visited)
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	moves := strings.Split(strings.TrimSpace(string(dat)), "\n")

	t1 := time.Now()
	fmt.Printf("First part: %d\n", simulate(moves, 2))
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	t2 := time.Now()
	fmt.Printf("Second part: %d\n", simulate(moves, 10))
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
}

func sign(n int) int {
	if n > 0 {
		return 1
	}
	return -1
}

func atoi(str string) int {
	n, e := strconv.Atoi(str)
	check(e)
	return n
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
