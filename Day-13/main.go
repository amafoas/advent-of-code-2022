package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	pairs := strings.Split(strings.TrimSpace(string(dat)), "\n\n")

	pcks, p1 := []any{}, 0
	for i, pck := range pairs {
		pck := strings.Split(pck, "\n")
		var left, right any
		json.Unmarshal([]byte(pck[0]), &left)
		json.Unmarshal([]byte(pck[1]), &right)
		pcks = append(pcks, left, right)

		if compare(left, right) <= 0 {
			p1 += i + 1
		}
	}

	fmt.Println("First part: ", p1)

	pcks = append(pcks, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(pcks, func(i, j int) bool { return compare(pcks[i], pcks[j]) < 0 })

	p2 := 1
	for i, pck := range pcks {
		if fp := fmt.Sprint(pck); fp == "[[2]]" || fp == "[[6]]" {
			p2 *= i + 1
		}
	}

	fmt.Println("Second part: ", p2)
}

func compare(left any, right any) int {
	lv, lftArr := left.([]any)
	rv, rgtArr := right.([]any)

	if !lftArr && !rgtArr { // both numbers
		return int(left.(float64) - right.(float64))
	}

	if !lftArr { // convert left to []
		lv = []any{left}
	}
	if !rgtArr { // convert right to []
		rv = []any{right}
	}

	for i := 0; i < len(lv) && i < len(rv); i++ {
		if c := compare(lv[i], rv[i]); c != 0 {
			return c
		}
	}

	return len(lv) - len(rv)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
