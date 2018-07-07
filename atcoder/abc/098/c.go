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
	n := intv()
	s := strv()
	wSum := make([]int, n)
	eSum := make([]int, n)
	if s[0] == 'W' {
		wSum[0] = 1
	}
	if s[n-1] == 'E' {
		eSum[n-1] = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == 'W' {
			wSum[i] = wSum[i-1] + 1
		} else {
			wSum[i] = wSum[i-1]
		}
	}
	for i := n - 2; i > -1; i-- {
		if s[i] == 'E' {
			eSum[i] = eSum[i+1] + 1
		} else {
			eSum[i] = eSum[i+1]
		}
	}
	ans := INF
	for i := 0; i < n; i++ {
		ans = min(ans, wSum[i]+eSum[i])
	}
	fmt.Println(ans - 1)
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
