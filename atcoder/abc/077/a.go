package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in        = bufio.NewReader(os.Stdin)
	out       = bufio.NewWriter(os.Stdout)
	MOD       = 1e9 + 7
	UMOD      = uint64(1e9 + 7)
	factorial []uint64
	inverse   []uint64
)

func main() {
	c := make([][]string, 2)
	copy := make([][]string, 2)
	for i, _ := range c {
		c[i] = make([]string, 3)
		copy[i] = make([]string, 3)
		row := strings.Split(strValue(), "")
		for j, v := range row {
			c[i][j] = v
		}
	}
	for i, v := range c {
		for j, _ := range v {
			copy[i][j] = c[1-i][2-j]
		}
	}
	for i, v := range c {
		for j, _ := range v {
			if c[i][j] != copy[i][j] {
				fmt.Println("NO")
				return
			}
		}
	}
	fmt.Println("YES")
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

func strValue() string {
	return strSlice()[0]
}

func strSlice() []string {
	line := strings.Split(readln(), " ")
	return line
}

func intValue() int {
	return intSlice()[0]
}

func intSlice() []int {
	line := strSlice()
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

func IntMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func IntMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
