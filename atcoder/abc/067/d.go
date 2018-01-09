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

type queue struct {
	node int
	path []int
}

func main() {
	n := intSlice()[0]
	adj := make([][]bool, n+1)
	for i, _ := range adj {
		adj[i] = make([]bool, n+1)
	}
	for i := 0; i < n-1; i++ {
		line := intSlice()
		a := line[0]
		b := line[1]
		adj[a][b] = true
		adj[b][a] = true
	}
	t1 := make([]int, n+1)
	t2 := make([]int, n+1)
	q := int{}
	for i := 0; i <= n; i++ {
		if adj[1][i] {
			q = append(q, i)
			t1[i] = 1
		}
		if adj[n][i] {
			t2[i] = 1
		}
	}
	for {
		if len(q) == 0 {
			break
		}
		now := q[0]
		q = q[1:]

		if now != n {

		}

	}

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
