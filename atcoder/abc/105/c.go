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
	bit []int
)

func main() {
	n := intv()
	if n == 0 {
		fmt.Println(0)
		return
	}
	b := make([]int, 32)
	for i, _ := range b {
		b[i] = int(math.Pow(float64(-2), float64(i)))
	}
	fmt.Println(b)
	_ = rec(n, b)
	/*
			ans := ""
			for i, _ := range bit {

			}
		fmt.Println(reverse(ans))
	*/
	fmt.Println(bit)
}

func rec(x int, b []int) (res int) {
	for i, _ := range b {
		if x == b[i] {
			bit = append(bit, b[i])
			return 0
		}
	}
	if x > 0 {
		idx := 0
		m := INF
		for i, _ := range b {
			diff := x + b[i]
			if diff < m {
				idx = i
				m = diff
			}
			if abs(b[i]) > x {
				break
			}
		}
		bit = append(bit, b[idx])
		fmt.Printf("x = %d, m = %d, b[idx] = %d\n", x, m, b[idx])
		return rec(m, b)
	} else {

	}
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
