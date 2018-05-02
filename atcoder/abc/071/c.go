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
	in        = bufio.NewReader(os.Stdin)
	out       = bufio.NewWriter(os.Stdout)
	MOD       = 1e9 + 7
	UMOD      = uint64(1e9 + 7)
	factorial []uint64
	inverse   []uint64
)

func main() {
	n := intv()
	a := ints()
	sort.Ints(a)
	res := 0
	for i := len(a) - 1; i >= 3; i-- {
		if a[i] == a[i-1] && a[i-1] == a[i-2] && a[i-2] == a[i-3] {
			res = max(res, a[i]*a[i])
		}
	}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		if v, ok := m[a[i]]; ok && v >= 2 {
			continue
		} else {
			m[a[i]]++
		}
	}
	c := make([]int, 0)
	for key, value := range m {
		if value >= 2 {
			c = append(c, key)
		}
	}
	sort.Ints(c)
	if len(c) >= 2 {
		res = max(res, c[len(c)-1]*c[len(c)-2])
	}
	fmt.Println(res)
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
