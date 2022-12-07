package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	fmt.Println("First part: ", findMarker(4, dat))
	fmt.Println("Second part: ", findMarker(14, dat))
}

func findMarker(n int, dat []byte) int {
	len := len(dat)
	for i := 0; i < len; i++ {
		if i+n < len {
			seq := dat[i : i+n]
			if allDifferent(seq) {
				return (i + n)
			}
		}
	}
	return 0
}

func allDifferent(seq []byte) bool {
	set := make(map[byte]bool)
	for _, v := range seq {
		if set[v] {
			return false
		}
		set[v] = true
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
