package main

import (
	"bufio"
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
	n := intv()
	adj := make([][]int, n)
	for i, _ := range adj {
		adj[i] = make([]int, n)
		for j := 0; j < n; j++ {
			adj[i][j] = 1 << 29
		}
	}
	for i := 0; i < n-1; i++ {
		row := ints()
		a, b, c := row[0], row[1], row[2]
		adj[a][b] = c
		adj[b][a] = c
	}
	adj = warshallFloyd(adj, n)
}

func warshallFloyd(adj [][]int, n int) [][]int {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				adj[i][j] = min(adj[i][j], adj[i][k]+adj[k][j])
			}
		}
	}
	return adj
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
