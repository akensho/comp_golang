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
	f   [][]int
	p   [][]int
)

func main() {
	n := intv()
	f = make([][]int, n)
	p = make([][]int, n)
	for i, _ := range f {
		f[i] = ints()
	}
	for i, _ := range p {
		p[i] = ints()
	}
	res := 0
	for i := 0; i < n; i++ {

		//		res += max(rec(i, 0, true), rec(i, 0, false))
	}
	fmt.Println(res)
}

func rec(x, idx int, do bool) int {
	if do {
		if idx == 9 {
			return f[x][9] * p[x][9]
		} else {
			return f[x][idx]*p[x][idx] + max(rec(x, idx+1, true), rec(x, idx+1, false))
		}
	} else {
		if idx == 9 {
			return 0
		} else {
			return max(rec(x, idx+1, true), rec(x, idx+1, false))
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
