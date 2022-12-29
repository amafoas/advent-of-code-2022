package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pos struct {
	x, y, z int
}

func add(p Pos, m Pos) Pos {
	return Pos{p.x + m.x, p.y + m.y, p.z + m.z}
}

var faces []Pos = []Pos{
	{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
	{-1, 0, 0}, {0, -1, 0}, {0, 0, -1},
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	cubes := make(map[Pos]bool)
	var limit int
	for _, cube := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		pos := getPosition(cube)
		limit = max(limit, max(pos.x, max(pos.y, pos.z)))
		cubes[pos] = true
	}

	t1 := time.Now()
	fmt.Printf("First part: %d\n", partOne(cubes))
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	t2 := time.Now()
	fmt.Printf("First part: %d\n", floodFill(Pos{0, 0, 0}, cubes, limit))
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
}

func partOne(cubes map[Pos]bool) int {
	var p1 int
	for cube := range cubes {
		for _, face := range faces {
			f := add(cube, face)
			if !cubes[f] {
				p1++
			}
		}
	}
	return p1
}

func floodFill(start Pos, cubes map[Pos]bool, limit int) int {
	stack := []Pos{start}
	visited := map[Pos]bool{start: true}

	var ext int
	for len(stack) > 0 {
		cb := stack[0]
		stack = stack[1:]

		/// adding 1 extra space to limits so code can propagate properly
		if cb.x < -1 || cb.y < -1 || cb.z < -1 ||
			cb.x > limit+1 || cb.y > limit+1 || cb.z > limit+1 {
			continue
		}

		for _, face := range faces {
			p := add(cb, face)
			if cubes[p] {
				ext++
			} else if !visited[p] {
				visited[p] = true
				stack = append(stack, p)
			}
		}
	}

	return ext
}

func getPosition(cube string) Pos {
	cords := strings.Split(cube, ",")
	x, _ := strconv.Atoi(cords[0])
	y, _ := strconv.Atoi(cords[1])
	z, _ := strconv.Atoi(cords[2])
	return Pos{x, y, z}
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
