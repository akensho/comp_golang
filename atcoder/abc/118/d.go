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

type orignal struct {
	value int
}

type orignals []orignal

func (o orignals) Len() int {
	return len(o)
}

func (o orignals) Less(i, j int) bool {
	return o[i].value < o[j].value
}

func (o orignals) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func main() {
	row := ints()
	n, _ := row[0], row[1]
	a := ints()
	b := make([]int, 0)
	for i, _ := range a {
		b = append(b, converter(a[i]))
	}
	sort.Ints(b)
	c := make([]int, 0)
	for i, _ := range b {
		c = append(c, reConverter(b[i]))
	}
	ans := ""
	for n > 0 {

	}
	fmt.Println(ans)
	/*
		for i, _ := range b {
			var_dump(reConverter(b[i]))
		}
		var_dump(n)
	*/
}

func getMatch(x int) int {
	if x == 1 {
		return 2
	}
	if x == 7 {
		return 3
	}
	if x == 4 {
		return 4
	}
	if x == 5 || x == 3 || x == 2 {
		return 5
	}
	if x == 9 || x == 6 {
		return 6
	}
	if x == 8 {
		return 7
	}
	return -1
}

func converter(x int) int {
	if x == 1 {
		return 0
	}
	if x == 7 {
		return 1
	}
	if x == 4 {
		return 2
	}
	if x == 5 {
		return 3
	}
	if x == 3 {
		return 4
	}
	if x == 2 {
		return 5
	}
	if x == 9 {
		return 6
	}
	if x == 6 {
		return 7
	}
	if x == 8 {
		return 8
	}
	return -1
}

func reConverter(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 7
	}
	if x == 2 {
		return 4
	}
	if x == 3 {
		return 5
	}
	if x == 4 {
		return 3
	}
	if x == 5 {
		return 2
	}
	if x == 6 {
		return 9
	}
	if x == 7 {
		return 6
	}
	if x == 8 {
		return 8
	}
	return -1
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
