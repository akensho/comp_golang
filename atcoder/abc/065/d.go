package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	MOD = 1e9 + 7
	es  = []Edge{}
)

type Edge struct {
	u, v, cost int
}

type Point struct {
	x, y int
}

func read() []Point {
	n := intSlice()[0]
	ps := make([]Point, n)
	for i := 0; i < n; i++ {
		row := intSlice()
		ps[i].x = row[0]
		ps[i].y = row[1]
	}
	return ps
}

func addEdge(n int, p []Point) {
	for i := 0; i < n-1; i++ {
		a := p[i].x
		b := p[i].y
		c := p[i+1].x
		d := p[i+1].y
		cost := IntMin(IntAbs(a-c), IntAbs(b-d))
		edge := Edge{
			u:    i,
			v:    i + 1,
			cost: cost,
		}
		es = append(es, edge)
	}
}

func kruskal(n int) int {
	var uf UnionFindTree
	sort.Slice(es, func(i, j int) bool {
		return es[i].cost < es[j].cost
	})
	for _, e := range es {
		fmt.Printf("u = %d, v = %d, cost = %d\n", e.u, e.v, e.cost)
	}
	uf.NewUnionFindTree(n)
	res := 0
	for _, e := range es {
		if !uf.IsSame(e.u, e.v) {
			uf.Unite(e.u, e.v)
			res += e.cost
		}
	}
	return res
}

func main() {
	ps := read()
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].x < ps[j].x
	})
	addEdge(len(ps), ps)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].y < ps[j].y
	})
	addEdge(len(ps), ps)
	res := kruskal(len(ps))
	fmt.Println(res)
}

/*
 * Convenient functions
 */

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

func IntMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func IntMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
 * Implementation of union find tree algorithm
 */

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
