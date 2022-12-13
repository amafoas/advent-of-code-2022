package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []uint64
	op        func(uint64) uint64
	test      func(uint64) int
	inspected uint64
	lcm       uint64
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	datSpl := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	monkeys := make([]Monkey, len(datSpl))

	for i, mky := range datSpl {
		monkeys[i] = newMonkey(strings.Split(mky, "\n"))
	}

	oneMonkeys := make([]Monkey, len(monkeys))
	copy(oneMonkeys, monkeys)
	fmt.Println("First part: ", monkeyBusiness(runMonkeys(oneMonkeys, 20, 3)))
	twoMonkeys := make([]Monkey, len(monkeys))
	copy(twoMonkeys, monkeys)
	fmt.Println("Second part: ", monkeyBusiness(runMonkeys(twoMonkeys, 10000, 1)))
}

func monkeyBusiness(monkeys []Monkey) uint64 {
	var a, b uint64 = 0, 0
	for _, mnk := range monkeys {
		ins := mnk.inspected
		if ins > a {
			b = a
			a = ins
		} else if ins > b {
			b = ins
		}
	}
	return a * b
}

func runMonkeys(monkeys []Monkey, rounds int, worryDivisor uint64) []Monkey {
	var lcm uint64 = 1
	for _, m := range monkeys {
		lcm *= m.lcm
	}

	for r := 0; r < rounds; r++ {
		for m := 0; m < len(monkeys); m++ {
			for _, item := range monkeys[m].items {
				worry := monkeys[m].op(item) / worryDivisor
				target := monkeys[m].test(worry)
				monkeys[target].items = append(monkeys[target].items, worry%lcm)
				monkeys[m].inspected++
			}
			monkeys[m].items = []uint64{}
		}
	}
	return monkeys
}

func newMonkey(mky []string) Monkey {
	monkey := Monkey{}
	rd, err := regexp.Compile(`\d+`)
	check(err)
	monkey.items = atoiArray(rd.FindAllString(mky[1], -1))

	r, err := regexp.Compile(`old [*-+] (\d+|old)`)
	check(err)
	monkey.op = parseOperation(strings.Split(r.FindString(mky[2]), " "))

	div := atoi(rd.FindString(mky[3]))
	t, f := int(atoi(rd.FindString(mky[4]))), int(atoi(rd.FindString(mky[5])))
	monkey.test = parseTest(div, t, f)

	monkey.inspected = 0
	monkey.lcm = div
	return monkey
}

func atoiArray(str []string) []uint64 {
	var res []uint64
	for _, c := range str {
		res = append(res, uint64(atoi(c)))
	}
	return res
}

func parseOperation(params []string) func(uint64) uint64 {
	if params[1] == "+" {
		return func(old uint64) uint64 { return (old + atoi(params[2])) }
	} else if params[1] == "*" && params[2] == "old" {
		return func(old uint64) uint64 { return (old * old) }
	}
	return func(old uint64) uint64 { return (old * atoi(params[2])) }
}

func parseTest(div uint64, t int, f int) func(uint64) int {
	return func(n uint64) int {
		if (n % div) == 0 {
			return t
		}
		return f
	}
}

func atoi(str string) uint64 {
	n, err := strconv.Atoi(str)
	check(err)
	return uint64(n)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
