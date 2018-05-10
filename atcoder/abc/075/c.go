package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in        = bufio.NewReader(os.Stdin)
	out       = bufio.NewWriter(os.Stdout)
	MOD       = 1e9 + 7
	UMOD      = uint64(1e9 + 7)
	factorial []uint64
	inverse   []uint64
)

func main() {
	row := ints()
	n, m := row[0], row[1]
	adj := make([][]bool, n)
	for i, _ := range adj {
		adj[i] = make([]bool, n)
	}
	type edge struct {
		from, to int
	}
	es := make([]edge, m)
	for i := 0; i < m; i++ {
		row = ints()
		a, b := row[0]-1, row[1]-1
		adj[a][b] = true
		adj[b][a] = true
		es[i].from = a
		es[i].to = b
	}
	ans := 0
	for _, e := range es {
		adj[e.from][e.to] = false
		adj[e.to][e.from] = false
		visited := make([]bool, n)
		visited = dfs(adj, visited, 0)
		for _, v := range visited {
			if !v {
				ans++
				break
			}
		}
		adj[e.from][e.to] = true
		adj[e.to][e.from] = true
	}
	fmt.Println(ans)
}

func dfs(adj [][]bool, visited []bool, x int) []bool {
	visited[x] = true
	for v := 0; v < len(visited); v++ {
		if !adj[x][v] {
			continue
		}
		if visited[v] {
			continue
		}
		visited = dfs(adj, visited, v)
	}
	return visited
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
