package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Elf struct {
	start int
	end   int
}

func newElf(s string) Elf {
	hours := strings.Split(s, "-")
	start, _ := strconv.Atoi(hours[0])
	end, _ := strconv.Atoi(hours[1])
	return Elf{start, end}
}

func (x Elf) ContainedIn(e Elf) bool {
	return e.start <= x.start && x.end <= e.end
}

func (x Elf) Overlaps(e Elf) bool {
	a := x.start <= e.start && e.start <= x.end
	b := x.start <= e.end && e.end <= x.end
	return x.ContainedIn(e) || e.ContainedIn(x) || a || b
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	pairs := strings.Split(strings.TrimSpace(string(dat)), "\n")

	t1 := time.Now()
	fully, overs := 0, 0
	for _, p := range pairs {
		v := strings.Split(p, ",")
		fElf := newElf(v[0])
		sElf := newElf(v[1])

		if fElf.ContainedIn(sElf) || sElf.ContainedIn(fElf) {
			fully++
			overs++
		} else if fElf.Overlaps(sElf) {
			overs++
		}
	}

	fmt.Printf("First part: %d\n", fully)
	fmt.Printf("Second part: %d\n", overs)
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
