package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	ind, val int
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	values := make(map[int]Pair) // og id as key
	var zero int
	var og []int
	for i, num := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		n, _ := strconv.Atoi(num)
		if n == 0 {
			zero = i
		}
		values[i] = Pair{i, n}
		og = append(og, n)
	}

	/// Part One
	fp := copyMap(values)
	mix(og, fp)
	fmt.Println("First part: ", sum(fp, zero))

	// Part Two
	sp := copyMap(values)
	mult(sp, 811589153)
	sog := parseToArray(values)
	for i := 0; i < 10; i++ {
		mix(sog, sp)
	}
	fmt.Println("Second part: ", sum(sp, zero))
}

func mix(og []int, values map[int]Pair) {
	for i := range og {
		p := values[i]
		ni := (p.ind + p.val) % (len(og) - 1)
		if ni < 0 {
			ni += (len(og) - 1)
		}
		values[i] = Pair{ni, p.val}
		for k, v := range values {
			if k != i {
				newp := v
				if p.ind < ni && p.ind < v.ind && v.ind <= ni {
					newp.ind--
				} else if ni < p.ind && ni <= v.ind && v.ind < p.ind {
					newp.ind++
				}
				values[k] = newp
			}
		}
	}
}

func mult(values map[int]Pair, dk int) {
	for k, v := range values {
		values[k] = Pair{v.ind, v.val * dk}
	}
}

func sum(values map[int]Pair, zero int) int {
	var sum int
	arr := parseToArray(values)
	for i := values[zero].ind + 1000; i <= values[zero].ind+3000; i += 1000 {
		idx := i % (len(values))
		sum += arr[idx]
	}
	return sum
}

func parseToArray(m map[int]Pair) []int {
	arr := make([]int, len(m))
	for _, v := range m {
		arr[v.ind] = v.val
	}
	return arr
}

func copyMap(m map[int]Pair) map[int]Pair {
	copy := make(map[int]Pair)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
