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
	s := strv()
	t := strv()
	ss := make([]string, 0)
	tt := make([]string, 0)
	for _, v := range s {
		ss = append(ss, string(v))
	}
	for _, v := range t {
		tt = append(tt, string(v))
	}
	sort.Strings(ss)
	sort.Strings(tt)
	fmt.Println(ss)
	fmt.Println(tt)
	conv := make(map[string]string)
	for i := 0; i < len(ss); i++ {
		if ss[i] == tt[i] {
			continue
		}
		if val, ok := conv[ss[i]]; ok {
			if val == tt[i] {
				break
			}
			fmt.Println("No")
			return
		}
		conv[ss[i]] = tt[i]
		conv[tt[i]] = ss[i]
	}
	//	fmt.Println(conv)
	fmt.Println("Yes")
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
