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
	s := strv()
	t := strv()
	conv := make(map[string]string)
	set := make(map[string]bool)
	for i, _ := range s {
		if s[i] == t[i] {
			continue
		}
		c1, c2 := string(s[i]), string(t[i])
		if value, ok := conv[c1]; ok {
			if c2 != value {
				fmt.Println("No")
				return
			}
		} else {
			if !set[c2] {
				conv[c1] = c2
				set[c2] = true
			} else {
				fmt.Println("No")
				return
			}
		}
	}
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
