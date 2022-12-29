package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	var elves []string = strings.Split(strings.TrimSpace(string(dat)), "\n\n")

	var calories []int
	for _, elf := range elves {
		e := strings.Split(elf, "\n")
		elfCalories := 0
		for _, cal := range e {
			v, err := strconv.Atoi(cal)
			check(err)
			elfCalories += v
		}
		calories = append(calories, elfCalories)
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	t1 := time.Now()
	fmt.Printf("First part: %d\n", calories[0])
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	t2 := time.Now()
	fmt.Printf("Second part: %d\n", calories[0]+calories[1]+calories[2])
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
