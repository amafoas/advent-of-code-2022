package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Pair struct {
	x int
	y int
}

type Sensor struct {
	pos      Pair
	beacon   Pair
	distance int
}

// calculates the manhattan distance
func distance(p Pair, m Pair) int {
	return abs(p.x-m.x) + abs(p.y-m.y)
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	r, _ := regexp.Compile(`-?\d+`)
	var xmin, xmax int = math.MaxInt, 0 // width limits
	var locations []Sensor
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		values := atoiAll(r.FindAllString(line, -1))
		var s, b Pair = Pair{values[0], values[1]}, Pair{values[2], values[3]}
		var d int = distance(s, b)
		if s.x-d < xmin {
			xmin = s.x - d
		}
		if s.x+d > xmax {
			xmax = s.x + d
		}
		locations = append(locations, Sensor{s, b, d})
	}

	t1 := time.Now()
	fmt.Printf("First part: %d\n", partOne(locations, xmin, xmax, 2000000))
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	t2 := time.Now()
	fmt.Printf("First part: %d\n", partTwo(locations, 4000000))
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
}

func partOne(locations []Sensor, xmin, xmax, ypos int) int {
	covered := 0
	for l := xmin; l <= xmax; l++ {
		position := Pair{l, ypos}
		for _, sensor := range locations {
			onRange := sensor.distance >= distance(sensor.pos, position)
			notSensor := position != sensor.pos
			notBeacon := position != sensor.beacon
			if onRange && notSensor && notBeacon {
				covered++
				break
			}
		}
	}
	return covered
}

func partTwo(locations []Sensor, limit int) int {
	for y := 0; y <= limit; y++ {
		for x := 0; x <= limit; x++ {
			position := Pair{x, y}
			notCovered := true
			for _, sensor := range locations {
				dm := distance(sensor.pos, position)
				onRange := sensor.distance >= dm
				if onRange {
					x += sensor.distance - dm
					notCovered = false
					break
				}
			}
			if notCovered {
				return (position.x * 4000000) + position.y
			}
		}
	}
	return -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func atoiAll(sarr []string) []int {
	var narr []int
	for i := range sarr {
		n, _ := strconv.Atoi(sarr[i])
		narr = append(narr, n)
	}
	return narr
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
