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
	s := strValue()
	t := strValue()
	ss := make([]rune, 0)
	tt := make([]rune, 0)
	for _, v := range s {
		ss = append(ss, v)
	}
	for _, v := range t {
		tt = append(tt, v)
	}
	sort.Sort(RuneSlice(ss))
	sort.Sort(RuneSlice(tt))
	s_ := string(ss)
	t_ := reverse(string(tt))
	if s_ < t_ {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

type RuneSlice []rune

func (r RuneSlice) Len() int {
	return len(r)
}

func (r RuneSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r RuneSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

func strValue() string {
	return strSlice()[0]
}

func strSlice() []string {
	line := strings.Split(readln(), " ")
	return line
}

func intValue() int {
	return intSlice()[0]
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
