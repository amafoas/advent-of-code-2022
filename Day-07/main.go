package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	parent *Node
	childs map[string]*Node
	size   int
	name   string
}

func node(parent *Node, name string, size int) *Node {
	return &Node{
		parent: parent,
		childs: make(map[string]*Node),
		size:   size,
		name:   name,
	}
}

func (n *Node) addNewNode(name string, size int) {
	m := node(n, name, size)
	n.childs[name] = m
	p := n
	for p != nil && size > 0 {
		p.size += size
		p = p.parent
	}
}

func (n *Node) isLeaf() bool {
	return len(n.childs) == 0
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)
	ins := strings.Split(string(dat), "\n")

	var curr *Node = node(nil, "root", 0)
	for i := 0; i < len(ins); i++ {
		line := strings.Split(ins[i], " ")
		if line[0] == "$" {
			if line[1] == "cd" {
				if line[2] == ".." {
					curr = curr.parent
				} else {
					curr.addNewNode(line[2], 0)
					curr = curr.childs[line[2]]
				}
			}
		} else if line[0] != "dir" {
			size, err := strconv.Atoi(line[0])
			check(err)
			curr.addNewNode(line[1], size)
		}
	}

	root := curr
	for root.parent != nil {
		root = root.parent
	}

	fmt.Println("First part: ", partOne(root, 100000))
	fmt.Println("Second part: ", partTwo(root, root.size-40000000))
}

func partOne(fs *Node, min int) int {
	var under int
	for _, child := range fs.childs {
		if !child.isLeaf() {
			if child.size <= min {
				under += child.size
			}
			under += partOne(child, min)
		}
	}
	return under
}

func partTwo(fs *Node, need int) int {
	var del int = math.MaxInt
	for _, child := range fs.childs {
		if !child.isLeaf() {
			if child.size >= need && del > child.size {
				del = child.size
			}
			del = min(partTwo(child, need), del)
		}
	}
	return del
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
