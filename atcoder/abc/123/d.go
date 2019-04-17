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
	x, y, _, k := row[0], row[1], row[2], row[3]
	a := ints()
	b := ints()
	c := ints()

	s1 := make([]int, 0)
	for i, _ := range a {
		for j, _ := range b {
			s1 = append(s1, a[i]+b[j])
		}
	}
	sort.Ints(c)
	sort.Ints(s1)
	s1 = reverse(s1)
	c = reverse(c)
	s1 = s1[:min(k, x*y)]
	//fmt.Println(s1)
	s2 := make([]int, 0)
	for i, _ := range s1 {
		for j, _ := range c {
			s2 = append(s2, s1[i]+c[j])
		}
	}
	sort.Ints(s2)
	s2 = reverse(s2)
	for i := 0; i < k; i++ {
		fmt.Println(s2[i])
	}
}

func reverse(a []int) []int {
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

func floatv() float64 {
	return floats()[0]
}

func floats() []float64 {
	line := strs()
	slice := make([]float64, 0)
	for _, tmp := range line {
		val, err := strconv.ParseFloat(tmp, 64)
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
