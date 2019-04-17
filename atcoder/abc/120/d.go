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
	INF = (1 << 32) - 1
)

func main() {
	row := ints()
	n, m := row[0], row[1]
	a := make([]int, m)
	b := make([]int, m)
	for i := m - 1; i > -1; i-- {
		row = ints()
		a[i], b[i] = row[0], row[1]
	}
	var uf ufTree
	uf.init(n + 1)
	ans := make([]int, m+1)
	ans[0] = n * (n - 1) / 2
	for i := 0; i < m; i++ {
		e1, e2 := a[i], b[i]
		if uf.same(e1, e2) {
			ans[i+1] = ans[i]
			continue
		}
		n1, n2 := uf.getSize(e1), uf.getSize(e2)
		ans[i+1] = ans[i] - n1*n2
		uf.unite(e1, e2)
	}
	for i := m - 1; i > -1; i-- {
		fmt.Println(ans[i])
	}
}

type ufTree struct {
	parent, rank, size []int
}

func (uf *ufTree) init(n int) {
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	uf.size = make([]int, n)
	for i, _ := range uf.parent {
		uf.parent[i] = i
	}
	for i, _ := range uf.rank {
		uf.rank[i] = 0
	}
	for i, _ := range uf.size {
		uf.size[i] = 1
	}
}

func (uf *ufTree) find(x int) int {
	if uf.parent[x] == x {
		return x
	}
	uf.parent[x] = uf.find(uf.parent[x])
	return uf.parent[x]
}

func (uf *ufTree) unite(x, y int) {
	px, py := uf.find(x), uf.find(y)

	if px == py {
		return
	}

	s := uf.size[px] + uf.size[py]

	if uf.rank[px] < uf.rank[py] {
		uf.parent[px] = py
		uf.size[py] = s
		return
	}
	uf.parent[py] = px
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
	uf.size[px] = s
}

func (uf *ufTree) same(x, y int) bool {
	if uf.find(x) == uf.find(y) {
		return true
	}
	return false
}

func (uf *ufTree) getSize(x int) int {
	px := uf.find(x)
	return uf.size[px]
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

func floatv() float64 {
	return floats()[0]
}

func floats() []float64 {
	line := strs()
	slice := make([]float64, 0)
	for _, tmp := range line {
		val, err := strconv.ParseFloat(tmp, 64)
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

func var_dump(value ...interface{}) {
	for _, v := range value {
		fmt.Printf("%#v\n", v)
	}
}
