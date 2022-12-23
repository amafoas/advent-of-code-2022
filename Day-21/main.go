package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	yelled bool
	val    int
	op     string
	left   string
	right  string
	caller string
}

func (m *Monkey) yell(n int) {
	m.yelled = true
	m.val = n
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	monkeys := make(map[string]*Monkey)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		l := strings.Split(line, " ")
		m := Monkey{}

		if len(l[1:]) == 1 {
			n, _ := strconv.Atoi(strings.TrimSpace(l[1]))
			m.yell(n)
		} else {
			m.left = l[1]
			m.op = l[2]
			m.right = l[3]
		}
		monkeys[l[0][:4]] = &m
	}

	fmt.Println("First part: ", DFS_Solve("root", monkeys))

	// Connect monkeys
	for k, v := range monkeys {
		if v.left != "" {
			monkeys[v.left].caller = k
		}
		if v.right != "" {
			monkeys[v.right].caller = k
		}
	}

	fmt.Println("Second part: ", partTwo(monkeys))
}

func partTwo(monkeys map[string]*Monkey) int {
	/// trace path from humn to root
	var path []string
	curr := monkeys["humn"].caller
	for curr != "root" {
		path = append(path, curr)
		curr = monkeys[curr].caller
	}

	var total int
	if m := monkeys["root"]; m.left != path[len(path)-1] {
		total = monkeys[m.left].val
	} else {
		total = monkeys[m.right].val
	}
	for i := len(path) - 1; i >= 0; i-- {
		m := monkeys[path[i]]
		if (i == 0 && m.left != "humn") || (i > 0 && m.left != path[i-1]) {
			total = inverse_eval(monkeys[m.left].val, total, m.op, true)
		} else {
			total = inverse_eval(monkeys[m.right].val, total, m.op, false)
		}
	}

	return total
}

// x op a = b   ==>   x = b rop a
func inverse_eval(a, b int, op string, left bool) int {
	var res int
	if op == "+" {
		res = b - a
	} else if op == "-" {
		if left { // a on left of -
			res = (b - a) / -1
		} else {
			res = b + a
		}
	} else if op == "*" {
		res = b / a
	} else if op == "/" {
		if left { // a on left of /
			res = a / b
		} else {
			res = b * a
		}
	}
	return res
}

func DFS_Solve(start string, monkeys map[string]*Monkey) int {
	stack := []string{start}
	waiting := []string{}
	visited := make(map[string]bool)

	for len(waiting) > 0 || len(stack) > 0 {
		for len(stack) > 0 {
			name := stack[0]
			stack = stack[1:]
			m := monkeys[name]

			if m.yelled {
				continue
			}

			l, r := monkeys[m.left], monkeys[m.right]
			if l.yelled && r.yelled {
				m.yell(eval(l.val, r.val, m.op))
			} else {
				if !visited[m.left] {
					visited[m.left] = true
					stack = append(stack, m.left)
				}
				if !visited[m.right] {
					visited[m.right] = true
					stack = append(stack, m.right)
				}
				waiting = append(waiting, name)
			}
		}
		stack = append(stack, waiting...)
		waiting = []string{}
	}

	return monkeys[start].val
}

func eval(a, b int, op string) int {
	var res int
	if op == "+" {
		res = a + b
	} else if op == "/" {
		res = a / b
	} else if op == "*" {
		res = a * b
	} else if op == "-" {
		res = a - b
	}
	return res
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
