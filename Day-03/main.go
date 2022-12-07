package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	rucksacks := strings.Split(strings.TrimSpace(string(dat)), "\n")

	fmt.Println("First part: ", partOne(rucksacks))
	fmt.Println("Second part: ", partTwo(rucksacks, len(rucksacks)))
}

func partOne(rucksacks []string) int {
	var res []byte
	for _, rucksack := range rucksacks {
		hl := len(rucksack) / 2
		l, r := makeSet([]byte(rucksack[:hl])), makeSet([]byte(rucksack[hl:]))

		res = append(res, getOcurrences([][]byte{l, r}, 2)...)
	}
	return calculate_points(res)
}

func partTwo(rucksacks []string, len int) int {
	var res []byte
	for i := 0; i < len; i += 3 {
		f := makeSet([]byte(rucksacks[i]))
		s := makeSet([]byte(rucksacks[i+1]))
		t := makeSet([]byte(rucksacks[i+2]))

		res = append(res, getOcurrences([][]byte{f, s, t}, 3)...)
	}
	return calculate_points(res)
}

func makeSet(arr []byte) []byte {
	var set []byte
	ocurred := make(map[byte]bool)
	for _, v := range arr {
		if !ocurred[v] {
			ocurred[v] = true
			set = append(set, v)
		}
	}
	return set
}

func getOcurrences(arr [][]byte, len int) []byte {
	ocurred := make(map[byte]int)
	var duppeds []byte
	for _, ar := range arr {
		for _, v := range ar {
			ocurred[v] += 1
			if ocurred[v] == len {
				duppeds = append(duppeds, v)
			}
		}
	}
	return duppeds
}

func calculate_points(arr []byte) int {
	var points int
	for _, v := range arr {
		if 65 <= v && v <= 90 {
			points += int(v) - 64 + 26
		} else if 97 <= v && v <= 122 {
			points += int(v) - 96
		}
	}
	return points
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
