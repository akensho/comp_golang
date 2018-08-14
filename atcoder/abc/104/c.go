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
	d, g := row[0], row[1]
	g = g / 100
	p := make([]int, d+1)
	c := make([]int, d+1)
	for i := 1; i <= d; i++ {
		row = ints()
		p[i], c[i] = row[0], row[1]
		c[i] = c[i] / 100
	}

	sum := 0
	for i := 1; i <= d; i++ {
		sum += p[i]*i + c[i]
	}
	dp := make([][]int, d+1)
	for i, _ := range dp {
		dp[i] = make([]int, sum+1)
	}

	fmt.Println(p)
	fmt.Println(c)
	fmt.Println(g)
	fmt.Println(dp)
	fmt.Println(len(dp))
	m := 2 >> 32
	for i := 0; i < len(dp)-1; i++ {
		if i == 0 {
			continue
		}
		for j := 0; j < len(dp[0])-1; j++ {
			if i+1 >= j {
				dp[i][j] = 1
				continue
			} else {
				dp[i+1][j] = min(dp[i][j], dp[i+1][j-1]+1)
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			m = min(m, dp[i][j])
		}
	}
	fmt.Println(dp)
	fmt.Println(m)
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
