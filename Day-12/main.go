package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Pair struct {
	x int
	y int
}

func pairSum(p Pair, m Pair) Pair {
	return Pair{p.x + m.x, p.y + m.y}
}

type Grid struct {
	values  [][]string
	len     Pair
	initial Pair
	end     Pair
}

func (g *Grid) onRange(p Pair) bool {
	onRangeX := 0 <= p.x && p.x < g.len.x
	onRangeY := 0 <= p.y && p.y < g.len.y
	return onRangeX && onRangeY
}

func (g *Grid) elevationBetween(from Pair, to Pair) int {
	curr := g.values[from.x][from.y]
	next := g.values[to.x][to.y]
	return int([]byte(next)[0]) - int([]byte(curr)[0])
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	var grid Grid = parseDataToGrid(string(dat))

	fmt.Println("First part: ", dijkstra(grid))
	fmt.Println("Second part: ", partTwo(grid))
}

func partTwo(g Grid) int {
	min := math.MaxInt
	for x, row := range g.values {
		for y, val := range row {
			if val == "a" {
				g.initial = Pair{x, y}
				dij := dijkstra(g)
				if dij > 0 && dij < min {
					min = dij
				}
			}
		}
	}

	return min
}

func parseDataToGrid(dat string) Grid {
	var grid Grid
	for x, line := range strings.Split(strings.TrimSpace(dat), "\n") {
		arr := strings.Split(string(line), "")
		for y, v := range arr {
			if v == "S" {
				arr[y] = "a"
				grid.initial = Pair{x, y}
			} else if v == "E" {
				arr[y] = "z"
				grid.end = Pair{x, y}
			}
		}
		grid.values = append(grid.values, arr)
	}
	grid.len = Pair{len(grid.values), len(grid.values[0])}
	return grid
}

func dijkstra(g Grid) int {
	queue := [][3]int{{g.initial.x, g.initial.y, 0}}
	visited := map[Pair]bool{}

	for len(queue) > 0 {
		front := queue[0]
		frontPair := Pair{front[0], front[1]}
		queue = queue[1:]
		if visited[frontPair] {
			continue
		}
		visited[frontPair] = true

		if frontPair == g.end {
			return front[2]
		}

		for _, move := range []Pair{{0, 1}, {-1, 0}, {0, -1}, {1, 0}} {
			next := pairSum(frontPair, move)
			if g.onRange(next) && g.elevationBetween(frontPair, next) <= 1 {
				queue = append(queue, [3]int{next.x, next.y, front[2] + 1})
			}
		}
	}

	return -1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
