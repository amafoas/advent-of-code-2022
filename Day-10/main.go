package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./input.txt")
	check(err)
	data := strings.Split(strings.TrimSpace(string(raw)), "\n")

	var monitor [6][40]string
	X, cycle, strengths := 1, 0, 0

	for _, line := range data {
		ins := strings.Split(line, " ")

		if ins[0] == "noop" {
			strengths += increaseClock(&cycle, 1, X, &monitor)
		} else {
			strengths += increaseClock(&cycle, 2, X, &monitor)
			X += atoi(ins[1])
		}
	}

	fmt.Println("First part: ", strengths)
	fmt.Println("Second part: ")
	for _, row := range monitor {
		s := ""
		for _, v := range row {
			s += v
		}
		fmt.Println(s)
	}
}

func increaseClock(c *int, n int, reg int, mon *[6][40]string) int {
	signalStr := 0
	for i := 0; i < n; i++ {
		drawPixel(*c, reg, mon)
		*c++
		if (*c % 40) == 20 {
			signalStr += ((*c) * reg)
		}
	}
	return signalStr
}

func drawPixel(c int, reg int, mon *[6][40]string) {
	f, m := c/40, c%40
	if m == reg-1 || m == reg || m == reg+1 {
		mon[f][m] = "#"
	} else {
		mon[f][m] = "."
	}
}

func atoi(str string) int {
	n, err := strconv.Atoi(str)
	check(err)
	return n
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
