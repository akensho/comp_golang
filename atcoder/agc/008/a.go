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
	row := ints()
	x, y := row[0], row[1]

	if abs(x) == abs(y) {
		fmt.Println(1)
		return
	}

	type queue struct {
		x, cnt int
	}

	q := make([]queue, 0)
	q = append(q, queue{x: x, cnt: 0})
	for {
		now := q[0]
		q = q[1:]
		nx := now.x + 1
		cnt := now.cnt + 1
		if nx == y {
			fmt.Println(cnt)
			break
		}
		q = append(q, queue{x: nx, cnt: cnt})
		nnx := -now.x
		if nnx == y {
			fmt.Println(cnt)
			break
		}
		q = append(q, queue{x: nnx, cnt: cnt})
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
