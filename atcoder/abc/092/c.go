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
	n := intv()
	a := make([]int, n+1)
	a[0] = 0
	tmp := ints()
	for i := 1; i < len(a); i++ {
		a[i] = tmp[i-1]
	}
	s := make([]int, n)
	s[0] = abs(a[0] - a[1])
	for i := 1; i < len(s); i++ {
		s[i] = s[i-1] + abs(a[i+1]-a[i])
	}
	fmt.Println(s)
	for i := 1; i <= n; i++ {
		cost := 0
		if i == n {
			cost = s[n-2] + abs(a[n-1]-a[0])
		} else {
			cost = s[i+1] + abs(a[i-1]-a[i+1]) + abs(a[n]-a[0])
		}
		fmt.Println(cost)
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
