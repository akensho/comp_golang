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
	from  int
	to    int
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

//
// o := make([]orignal, 0)
// sort.Sort(orignals(o))
//

func main() {
	row := ints()
	n, m := row[0], row[1]
	x := ints()
	if m == 1 {
		fmt.Println(0)
		return
	}
	if n >= m {
		fmt.Println(0)
		return
	}
	if n == 1 {
		ans := 0
		for _, val := range x {
			ans += abs(val)
		}
		fmt.Println(ans)
		return
	}
	sort.Ints(x)
	y := make([]orignal, m)
	for i, _ := range x {
		if i == 0 {
			y[0].value = abs(x[1] - x[0])
			y[0].from = 0
			y[0].to = 1
			continue
		}
		if i == m-1 {
			y[m-1].value = abs(x[m-1] - x[m-2])
			y[m-1].from = m - 1
			y[m-1].to = m - 2
			break
		}
		if abs(x[i]-x[i-1]) < abs(x[i]-x[i+1]) {
			y[i].value = abs(x[i] - x[i-1])
			y[i].from = i
			y[i].to = i - 1
		} else {
			y[i].value = abs(x[i] - x[i+1])
			y[i].from = i
			y[i].to = i + 1
		}
	}
	sort.Sort(orignals(y))
	y = y[:len(y)-(n-1)]
	z := make([]bool, m)
	ans := 0
	for _, point := range y {
		if z[point.from] && z[point.to] {
			continue
		}
		ans += point.value
		z[point.from] = true
		z[point.to] = true
	}
	fmt.Println(ans)
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
