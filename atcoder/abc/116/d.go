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
	row := ints()
	n, k := row[0], row[1]
	t := make([]int, n+1)
	d := make([]int, n+1)
	sushi := make(map[int][]int)
	for i := 1; i <= n; i++ {
		line := ints()
		t[i], d[i] = line[0], line[1]
		tmp := sushi[t[i]]
		tmp = append(tmp, d[i])
		sushi[t[i]] = tmp
	}
	for i, slice := range sushi {
		sort.Ints(slice)
		sushi[i] = slice
	}
	fmt.Println(sushi)
	fmt.Println(k)
}

func intReverse(a []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(a); i++ {
		res = append(res, a[len(a)-i-1])
	}
	return res
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

func var_dump(value ...interface{}) {
	for _, v := range value {
		fmt.Printf("%#v\n", v)
	}
}
