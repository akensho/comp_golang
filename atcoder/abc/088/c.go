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
	c := make([][]int, 3)
	for i, _ := range c {
		c[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		row := ints()
		for j := 0; j < 3; j++ {
			c[i][j] = row[j]
		}
	}

	for a0 := -10000; a0 <= 10000; a0++ {
		b := make([]int, 0)
		for j := 0; j < 3; j++ {
			b = append(b, c[0][j]-a0)
		}
		if check(a0, b, c) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

func check(a int, b []int, c [][]int) bool {
	if c[1][0]-b[0] == c[1][1]-b[1] && c[1][1]-b[1] == c[1][2]-b[2] {
		if c[2][0]-b[0] == c[2][1]-b[1] && c[2][1]-b[1] == c[2][2]-b[2] {
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
