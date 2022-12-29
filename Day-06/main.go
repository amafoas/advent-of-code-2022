package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	t1 := time.Now()
	fmt.Printf("First part: %d\n", findMarker(4, dat))
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	t2 := time.Now()
	fmt.Printf("Second part: %d\n", findMarker(14, dat))
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
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
