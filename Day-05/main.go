package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Stack struct {
	values []string
}

func (s *Stack) append(values string) {
	s.values = append(s.values, values)
}

func (s *Stack) push(values []string) {
	s.values = append(values, s.values...)
}

func (s *Stack) pop(n int) ([]string, bool) {
	if !(len(s.values) >= n) {
		return []string{}, false
	}
	values := []string{}
	values = append(values, s.values[:n]...)
	s.values = s.values[n:]
	return values, true
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	split := strings.Split(strings.TrimSpace(string(dat)), "\n\n")

	layout := strings.Split(split[0], "\n")
	stacksLen := ((len(layout[0]) + 1) / 4)
	stacks := make([]Stack, stacksLen)

	r, err := regexp.Compile("[A-Z]")
	check(err)
	for _, l := range layout {
		l += " " // need it for slice
		stack, start, end := 0, 0, 4
		for end <= len(l) {
			crate := r.FindString(l[start:end])
			if crate != "" {
				stacks[stack%stacksLen].append(crate)
			}

			stack++
			start = end
			end += 4
		}
	}

	instructions := strings.Split(strings.TrimSpace(split[1]), "\n")
	r, err = regexp.Compile(`\d+`)
	check(err)

	t1 := time.Now()
	f := make([]Stack, stacksLen)
	copy(f, stacks)
	fmt.Printf("First part: %s\n", partOne(instructions, f, r))
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	t2 := time.Now()
	s := make([]Stack, stacksLen)
	copy(s, stacks)
	fmt.Printf("Second part: %s\n", partTwo(instructions, s, r))
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
}

func partOne(instructions []string, stacks []Stack, r *regexp.Regexp) string {
	for _, line := range instructions {
		ins := r.FindAllString(line, -1)
		move, from, to := atoi(ins[0]), atoi(ins[1])-1, atoi(ins[2])-1
		pop, _ := stacks[from].pop(move)
		stacks[to].push(reverse(pop))
	}
	return getFirst(stacks)
}

func reverse(s []string) []string {
	l := len(s)
	r := make([]string, l)
	for i, v := range s {
		r[l-i-1] = v
	}
	return r
}

func getFirst(stacks []Stack) string {
	var res string
	for _, stack := range stacks {
		if arr, ok := stack.pop(1); ok {
			res += arr[0]
		}
	}
	return res
}

func atoi(s string) int {
	val, err := strconv.Atoi(s)
	check(err)
	return val
}

func partTwo(instructions []string, stacks []Stack, r *regexp.Regexp) string {
	for _, line := range instructions {
		ins := r.FindAllString(line, -1)
		move, from, to := atoi(ins[0]), atoi(ins[1])-1, atoi(ins[2])-1
		pop, _ := stacks[from].pop(move)
		stacks[to].push(pop)
	}
	return getFirst(stacks)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
