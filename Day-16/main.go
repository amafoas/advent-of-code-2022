package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Valve struct {
	flow    int
	tunnels []string
}

type Path struct {
	flow    int
	visited []string
}

var valves map[string]Valve
var distances map[string]map[string]int
var non_zero []string

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	rv, err := regexp.Compile(`[A-Z]{2}`)
	check(err)
	rf, err := regexp.Compile(`\d+`)
	check(err)

	valves = make(map[string]Valve)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		valve := rv.FindAllString(line, -1)
		flow, _ := strconv.Atoi(rf.FindString(line))
		v := Valve{flow, valve[1:]}
		valves[valve[0]] = v
		if flow > 0 {
			non_zero = append(non_zero, valve[0])
		}
	}

	distances = floydWarshall(valves)

	fmt.Println("First part: ", DFS("AA", 30, Path{0, []string{}}, make(map[string]bool))[0].flow)
}

func DFS(current string, time int, path Path, visited map[string]bool) []Path {
	visited[current] = true
	paths := []Path{path}
	for _, next := range non_zero {
		newTime := time - distances[current][next] - 1
		if visited[next] || next == current || newTime <= 0 {
			continue
		}
		newPath := Path{path.flow + (newTime * valves[next].flow), append(path.visited, current)}
		paths = append(paths, DFS(next, newTime, newPath, copyMap(visited))...)
	}

	sort.Slice(paths, func(i, j int) bool { return paths[i].flow > paths[j].flow })
	return paths
}

func copyMap(m map[string]bool) map[string]bool {
	mcopy := make(map[string]bool)
	for k, v := range m {
		mcopy[k] = v
	}
	return mcopy
}

func floydWarshall(valves map[string]Valve) map[string]map[string]int {
	var dist map[string]map[string]int = make(map[string]map[string]int)

	for i := range valves {
		for j := range valves {
			if _, ok := dist[i]; !ok {
				dist[i] = make(map[string]int)
			}
			if i == j {
				dist[i][j] = 0
			} else if contains(valves[i].tunnels, j) {
				dist[i][j] = 1
			} else {
				dist[i][j] = 999999
			}
		}
	}

	for k := range valves {
		for i := range valves {
			for j := range valves {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	return dist
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
