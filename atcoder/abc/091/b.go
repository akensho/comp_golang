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
	s := make([]string, n)
	for i, _ := range s {
		s[i] = strv()
	}
	m := intv()
	t := make([]string, m)
	for i, _ := range t {
		t[i] = strv()
	}
	tmp := map[string]int{}
	for _, v := range s {
		tmp[v]++
	}
	for _, v := range t {
		tmp[v]--
	}
	ans := -1
	for _, v := range tmp {
		ans = max(ans, v)
	}
	if ans < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(ans)
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
