package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	var forest [][]int
	for _, row := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		arr := make([]int, len(row))
		for i, char := range strings.Split(row, "") {
			num, _ := strconv.Atoi(char)
			arr[i] = num
		}
		forest = append(forest, arr)
	}

	t1 := time.Now()
	maxScore, visibles := 0, 0
	for y, row := range forest {
		for x := range row {
			isVisible, score := visibleAndScore(x, y, forest)
			if isVisible {
				visibles++
			}
			maxScore = max(maxScore, score)
		}
	}

	fmt.Println("First part: ", visibles)
	fmt.Println("Second part: ", maxScore)
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))
}

func visibleAndScore(x int, y int, grid [][]int) (bool, int) {
	coverL, coverR, coverD, coverU := false, false, false, false
	scoreL, scoreR, scoreD, scoreU := 0, 0, 0, 0
	for i := 1; i < len(grid) || i < len(grid[0]); i++ {
		if r := x + i; !coverR && r < len(grid[0]) {
			scoreR++
			coverR = grid[y][r] >= grid[y][x]
		}
		if l := x - i; !coverL && l >= 0 {
			scoreL++
			coverL = grid[y][l] >= grid[y][x]
		}
		if u := y + i; !coverU && u < len(grid) {
			scoreU++
			coverU = grid[u][x] >= grid[y][x]
		}
		if d := y - i; !coverD && d >= 0 {
			scoreD++
			coverD = grid[d][x] >= grid[y][x]
		}
	}

	isVisible := !(coverL && coverR && coverD && coverU)
	score := scoreL * scoreR * scoreD * scoreU
	return isVisible, score
}

func max(a int, b int) int {
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
