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
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	INF = (1 << 32) - 1
)

func main() {
	n := intv()
	l := []int{1}
	for i := 6; i <= 1000000; i *= 6 {
		l = append(l, i)
	}
	for i := 9; i <= 1000000; i *= 9 {
		l = append(l, i)
	}
	sort.Ints(l)
	x := 0
	for {
		idx := -1
		if l[len(l)-1] < n {
			idx = len(l) - 1
		} else {
			for i, v := range l {
				if n <= v {
					idx = i
					break
				}
			}
			if n == l[idx] {
				x++
				break
			}
		}
		x += n / l[idx-1]
		n = n % l[idx-1]
	}
	fmt.Println(x)
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
