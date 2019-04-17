package main

import (
	"bufio"
	"fmt"
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
	_, q := row[0], row[1]
	s := strv()
	l := make([]int, q)
	r := make([]int, q)
	for i := 0; i < q; i++ {
		row := ints()
		l[i], r[i] = row[0], row[1]
	}
	idx := 0
	sum := make([]int, len(s)+1)
	for idx < len(s) {
		now := string(s[idx])
		if now != "A" {
			idx++
			sum[idx] = sum[idx-1]
			continue
		}
		if idx == len(s)-1 {
			break
		}
		next := string(s[idx+1])
		if next == "C" {
			sum[idx+1] = sum[idx] + 1
		} else {
			sum[idx+1] = sum[idx]
		}
		idx++
	}
	//	fmt.Println(sum)
	for i := 0; i < q; i++ {
		fmt.Println(sum[r[i]-1] - sum[l[i]-1])
	}
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
