package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Grid struct {
	cells [][]string
	minW  int
	maxW  int
	maxH  int
}

func (g *Grid) get(y int, x int) (string, bool) {
	onLimitsX := g.minW <= x && x <= g.maxW
	onLimitsY := 0 <= y && y <= g.maxH
	if !(onLimitsX && onLimitsY) {
		return "", false
	}
	val := g.cells[y][x-g.minW]
	return val, true
}

func (g *Grid) edit(y int, x int, val string) {
	g.cells[y][x-g.minW] = val
}

func (grid *Grid) drawWallsFrom(input []string) {
	for _, line := range input {
		spl := strings.Split(line, " -> ")
		for i := 1; i < len(spl); i++ {
			x0, y0 := extracValues(spl[i-1])
			x1, y1 := extracValues(spl[i])
			for (x0 - x1) != 0 {
				grid.edit(y0, x0, "#")
				x0 += signum(x0-x1) * -1
			}
			for (y0 - y1) != 0 {
				grid.edit(y0, x0, "#")
				y0 += signum(y0-y1) * -1
			}
			grid.edit(y0, x0, "#")
		}
	}
}

func newGrid(minW, maxW, maxH int) Grid {
	cells := make([][]string, maxH+1)
	for i := 0; i < len(cells); i++ {
		s := make([]string, maxW-minW+1)
		for j := range s {
			s[j] = "."
		}
		cells[i] = s
	}

	return Grid{cells, minW, maxW, maxH}
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	input := strings.Split(strings.TrimSpace(string(dat)), "\n")

	var maxH, maxW, minW int = getLimits(input)

	var grid Grid = newGrid(minW, maxW, maxH)
	grid.drawWallsFrom(input)

	// Second part grid
	newH := maxH + 2
	var floorGrid Grid = newGrid((minW - newH), (maxW + newH), newH)
	var s string = fmt.Sprintf("%d,%d -> %d,%d", (minW - newH), newH, (maxW + newH), newH)
	floorGrid.drawWallsFrom(append(input, s))

	t1 := time.Now()
	fmt.Printf("First part: %d\n", simulateSand(&grid, 500, 0))
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	t2 := time.Now()
	fmt.Printf("First part: %d\n", simulateSand(&floorGrid, 500, 0))
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
}

func simulateSand(g *Grid, sandX int, sandY int) int {
	var sand int
	for {
		x, y := sandX, sandY
		sand++
		for {
			if val, ok := g.get(y, x); !ok {
				return sand - 1 // out of range
			} else if val == "o" || val == "#" {
				if l, lok := g.get(y, x-1); !lok {
					return sand - 1 // out of range
				} else if l == "." {
					x--
				} else if r, rok := g.get(y, x+1); !rok {
					return sand - 1 // out of range
				} else if r == "." {
					x++
				} else if y-1 == sandY {
					return sand // sink covered
				} else {
					g.edit(y-1, x, "o")
					break
				}
			}
			y++
		}
	}
}

func getLimits(input []string) (int, int, int) {
	var maxH, maxW, minW int = 0, 0, math.MaxInt
	for _, line := range input {
		for _, c := range strings.Split(line, " -> ") {
			x, y := extracValues(c)
			if y > maxH {
				maxH = y
			}
			if x < minW {
				minW = x
			}
			if x > maxW {
				maxW = x
			}
		}
	}
	return maxH, maxW, minW
}

func signum(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}

func extracValues(s string) (int, int) {
	cords := strings.Split(s, ",")
	x, _ := strconv.Atoi(cords[0])
	y, _ := strconv.Atoi(cords[1])
	return x, y
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
