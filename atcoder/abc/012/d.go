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
	n, m := row[0], row[1]
	matrix := make([][]int, n+1)
	for i, _ := range matrix {
		matrix[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			matrix[i][j] = 1 << 30
		}
	}
	for i := 0; i < m; i++ {
		row := ints()
		a, b, t := row[0], row[1], row[2]
		matrix[a][b] = t
		matrix[b][a] = t
	}
	for k := 0; k < n+1; k++ {
		for i := 0; i < n+1; i++ {
			for j := 0; j < n+1; j++ {
				matrix[i][j] = min(matrix[i][k]+matrix[k][j], matrix[i][j])
			}
		}
	}
	res := 1 << 30
	for i := 1; i < n+1; i++ {
		tmp := 0
		for j := 1; j < n+1; j++ {
			if i == j {
				continue
			}
			tmp = max(tmp, matrix[i][j])
		}
		res = min(tmp, res)
	}
	fmt.Println(res)
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
