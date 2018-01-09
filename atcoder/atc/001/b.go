package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type Query struct {
	p, a, b int
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

func main() {
	n, q := read()
	var uf UnionFindTree
	uf.NewUnionFindTree(n)
	for _, query := range q {
		if query.p == 1 {
			if uf.IsSame(query.a, query.b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		} else {
			uf.Unite(query.a, query.b)
		}
	}
}

func read() (int, []Query) {
	row := intSlice()
	n := row[0]
	q := row[1]
	query := make([]Query, q)
	for i, _ := range query {
		line := intSlice()
		tmp := Query{
			p: line[0],
			a: line[1],
			b: line[2],
		}
		query[i] = tmp
	}
	return n, query
}

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

func strSlice() []string {
	line := strings.Split(readln(), " ")
	return line
}

func intSlice() []int {
	line := strSlice()
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
