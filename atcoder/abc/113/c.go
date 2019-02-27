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

type city struct {
	year int
	idx  int
}

type cities []city

func (o cities) Len() int {
	return len(o)
}

func (o cities) Less(i, j int) bool {
	return o[i].year < o[j].year
}

func (o cities) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func main() {
	row := ints()
	n, m := row[0], row[1]
	c := make([][]city, n)
	for i, _ := range c {
		c[i] = make([]city, 0)
	}
	for i := 0; i < m; i++ {
		row := ints()
		num := row[0] - 1
		c[num] = append(c[num], city{idx: i, year: row[1]})
	}
	for i, _ := range c {
		sort.Sort(cities(c[i]))
	}
	//var_dump(c)
	ans := make([]string, m)
	for idx, val := range c {
		for i, v := range val {
			ans[v.idx] = padding(idx+1) + padding(i+1)
		}
	}
	for _, v := range ans {
		fmt.Println(v)
	}
}

func padding(x int) string {
	s := strconv.Itoa(x)
	for {
		if len(s) >= 6 {
			break
		}
		s = "0" + s
	}
	return s
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
