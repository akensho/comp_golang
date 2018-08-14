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
)

func main() {
	n := intv()
	if n == 1 || n == 2 {
		fmt.Println(0)
		return
	} else if n == 3 {
		fmt.Println(1)
		return
	}
	a := make([]int, n+1)
	a[0] = 0
	a[1] = 0
	a[2] = 0
	a[3] = 1
	for i := 4; i <= n; i++ {
		a[i] = (a[i-1] + a[i-2] + a[i-3]) % 10007
	}
	fmt.Println(a[n])
}

func contains(x int) bool {
	a := strconv.Itoa(x)
	for i, _ := range a {
		if a[i] == '3' {
			return true
		}
	}
	return false
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
