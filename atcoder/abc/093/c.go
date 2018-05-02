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
	row := ints()
	sort.Ints(row)
	a, b, c := row[0], row[1], row[2]
	cnt := 0
	for {
		if b == c {
			break
		}
		a++
		b++
		cnt++
	}
	diff := c - a
	if diff == 0 {
		fmt.Println(cnt)
		return
	}
	if diff == 1 {
		cnt += 2
		fmt.Println(cnt)
		return
	}
	//	fmt.Println(cnt)
	//	fmt.Println(a)
	//	fmt.Println(b)
	//	fmt.Println(c)
	if diff%2 == 0 {
		cnt += diff / 2
	} else {
		cnt += diff/2 + 2
	}
	fmt.Println(cnt)
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
