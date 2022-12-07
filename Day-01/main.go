package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	fmt.Print("First part: ", calories[0], "\n")
	fmt.Print("Second part: ", calories[0]+calories[1]+calories[2], "\n")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
