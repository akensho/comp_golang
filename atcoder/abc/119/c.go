package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	n, a, b, c := row[0], row[1], row[2], row[3]
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = intv()
	}
	m := int(math.Pow(4, float64(n)))
	pattern := make([]string, 0)
	for i := 0; i < m; i++ {
		tmp := strconv.FormatInt(int64(i), 4)
		pattern = append(pattern, padding(tmp, n))
	}
	ans := INF
	for _, v := range pattern {
		if check(v) {
			ans = min(ans, calc(v, n, a, b, c, l))
		}
	}
	fmt.Println(ans)
}

func check(s string) bool {
	f := make(map[string]bool)
	f["a"] = false
	f["b"] = false
	f["c"] = false

	for _, v := range s {
		switch v {
		case '0':
			continue
		case '1':
			f["a"] = true
		case '2':
			f["b"] = true
		case '3':
			f["c"] = true
		}
	}

	if a, _ := f["a"]; a {
		if b, _ := f["b"]; b {
			if c, _ := f["c"]; c {
				return true
			}
		}
	}
	return false
}

func calc(s string, n, a, b, c int, l []int) int {
	info := make([][]int, 4)
	for i, _ := range info {
		info[i] = make([]int, 0)
	}
	for i, v := range s {
		switch v {
		case '0':
			info[0] = append(info[0], i)
		case '1':
			info[1] = append(info[1], i)
		case '2':
			info[2] = append(info[2], i)
		case '3':
			info[3] = append(info[3], i)
		}
	}
	cost := 0
	for i, _ := range info {
		switch i {
		case 0:
			continue
		case 1:
			cost += calcCost(a, info[1], l)
		case 2:
			cost += calcCost(b, info[2], l)
		case 3:
			cost += calcCost(c, info[3], l)
		}
	}
	return cost
}

func calcCost(x int, info, l []int) int {
	bamboos := make([]int, len(info))
	for i, _ := range info {
		bamboos[i] = l[info[i]]
		if bamboos[i] == x {
			return 0
		}
	}
	cost := INF
	sum := 0
	for _, v := range bamboos {
		sum += v
	}
	for i, _ := range bamboos {
		cost = min(abs(bamboos[i]-x), cost)
	}
	cost = min(cost, abs(x-sum)+(len(bamboos)-1)*10)
	return cost
}

func padding(s string, x int) string {
	for {
		if len(s) >= x {
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
