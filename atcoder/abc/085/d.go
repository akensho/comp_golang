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
	n, h := row[0], row[1]
	s := make([]sord, n)
	for i := 0; i < n; i++ {
		row := ints()
		s[i].a = row[0]
		s[i].b = row[1]

	}
	sort.Sort(sords(s))
	bb := make([]b, 0)
	for i, v := range s {
		if v.b > s[0].a {
			bb = append(bb, b{idx: i, value: v.b})
		}
	}
	sort.Sort(bs(bb))
	cnt := 0
	for _, v := range bb {
		cnt++
		h = h - v.value
		if h <= 0 {
			fmt.Println(cnt)
			return
		}
	}
	for h > 0 {
		cnt++
		h = h - s[0].a
	}
	fmt.Println(cnt)
}

type b struct {
	idx, value int
}

type bs []b

func (o bs) Len() int {
	return len(o)
}

func (o bs) Less(i, j int) bool {
	return o[i].value > o[j].value
}

func (o bs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

type sord struct {
	a, b int
}

type sords []sord

func (o sords) Len() int {
	return len(o)
}

func (o sords) Less(i, j int) bool {
	return o[i].a > o[j].a
}

func (o sords) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
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
