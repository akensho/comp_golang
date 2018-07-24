package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	INF = (1 << 32) - 1
)

func main() {
	row := ints()
	n, m := row[0], row[1]

	var uf UnionFindTree
	uf.NewUnionFindTree(n)
}

type Edge struct {
	u, v, cost int
}

type UnionFindTree struct {
	Parent, Rank []int
}

func (uf *UnionFindTree) NewUnionFindTree(n int) {
	uf.Parent = make([]int, n)
	uf.Rank = make([]int, n)
	for i, _ := range uf.Parent {
		uf.Parent[i] = i
	}
	for i, _ := range uf.Rank {
		uf.Rank[i] = 0
	}
}

func (uf *UnionFindTree) Find(x int) int {
	if uf.Parent[x] == x {
		return x
	}
	uf.Parent[x] = uf.Find(uf.Parent[x])
	return uf.Parent[x]
}

func (uf *UnionFindTree) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if uf.Rank[x] < uf.Rank[y] {
		uf.Parent[x] = y
	} else {
		uf.Parent[y] = x
		if uf.Rank[x] == uf.Rank[y] {
			uf.Rank[x]++
		}
	}
}

func (uf *UnionFindTree) IsSame(x, y int) bool {
	if uf.Find(x) == uf.Find(y) {
		return true
	}
	return false
}

/* template functions */

func readln() string {
	buf := make([]byte, 0)
	for {
		line, prefix, err := in.ReadLine()
		if err != nil {
			panic(err)
		}
		buf = append(buf, line...)
		if prefix == false {
			break
		}
	}
	return string(buf)
}

func strv() string {
	return strs()[0]
}

func strs() []string {
	line := strings.Split(readln(), " ")
	return line
}

func intv() int {
	return ints()[0]
}

func ints() []int {
	line := strs()
	slice := make([]int, 0)
	for _, tmp := range line {
		val, err := strconv.Atoi(tmp)
		if err != nil {
			panic(err)
		}
		slice = append(slice, val)
	}
	return slice
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
