package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in   = bufio.NewReader(os.Stdin)
	out  = bufio.NewWriter(os.Stdout)
	INF  = (1 << 32) - 1
	n, m int
	g    [][]int
	set  map[edge]struct{}
)

type edge struct {
	u, v int
}

func main() {
	// read
	row := ints()
	n, m = row[0], row[1]
	g = make([][]int, n+1)
	for i, _ := range g {
		g[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			g[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		tmp := ints()
		a, b, c := tmp[0], tmp[1], tmp[2]
		g[a][b] = c
		g[b][a] = c
	}

	// solve
	set = make(map[edge]struct{}, 0)
	for i := 1; i <= n; i++ {
		prev := dijkstra(i)
		for j := 1; j <= n; j++ {
			if i == j {
				continue
			}
			path := getPath(j, prev)
			for idx := 0; idx < len(path)-1; idx++ {
				from := path[idx]
				to := path[idx+1]
				if _, ok := set[edge{u: to, v: from}]; !ok {
					set[edge{u: from, v: to}] = struct{}{}
				}
			}
		}
	}
	fmt.Println(m - len(set))
}

func reverse(a []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(a); i++ {
		res = append(res, a[len(a)-i-1])
	}
	return res
}

func getPath(t int, prev []int) (res []int) {
	for {
		if t == -1 {
			break
		}
		res = append(res, t)
		t = prev[t]
	}
	res = reverse(res)
	return res
}

func dijkstra(s int) []int {
	// init
	visited := make([]bool, n+1)
	for i, _ := range visited {
		visited[i] = false
	}
	dist := make([]int, n+1)
	for i, _ := range dist {
		dist[i] = INF
	}
	dist[s] = 0
	prev := make([]int, n+1)
	for i, _ := range prev {
		prev[i] = -1
	}

	// core
	for {
		v := -1
		for u := 0; u <= n; u++ {
			if !visited[u] {
				if v == -1 || dist[u] < dist[v] {
					v = u
				}
			}
		}
		if v == -1 {
			break
		}
		visited[v] = true

		for u := 0; u <= n; u++ {
			if dist[u] > dist[v]+g[v][u] {
				dist[u] = dist[v] + g[v][u]
				prev[u] = v
			}
		}
	}
	return prev
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
