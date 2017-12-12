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
	MIN = 1 << 29
)

func main() {
	n, m, a, b, c, d := read()
	res := make([]int, n)

	for i, _ := range res {
		min := MIN
		for j := 0; j < m; j++ {
			d := IntAbs(a[i]-c[j]) + IntAbs(b[i]-d[j])
			if d < min {
				res[i] = j + 1 // 1-index
				min = d
			}
		}
	}

	for _, ans := range res {
		fmt.Println(ans)
	}
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

func read() (n, m int, a, b, c, d []int) {
	l := intSlice()
	n = l[0]
	m = l[1]

	a = make([]int, n)
	b = make([]int, n)
	c = make([]int, m)
	d = make([]int, m)

	for i := 0; i < n; i++ {
		l = intSlice()
		a[i] = l[0]
		b[i] = l[1]
	}

	for i := 0; i < m; i++ {
		l = intSlice()
		c[i] = l[0]
		d[i] = l[1]
	}

	return n, m, a, b, c, d
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
