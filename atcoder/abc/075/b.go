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
	h, w := row[0], row[1]
	s := make([][]string, h)
	cp := make([][]string, h)
	for i, _ := range s {
		s[i] = make([]string, w)
		cp[i] = make([]string, w)
		r := strings.Split(strv(), "")
		for j, _ := range r {
			s[i][j] = r[j]
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "." {
				cp[i][j] = counter(s, i, j, h, w)
			} else {
				cp[i][j] = "#"
			}
		}
	}
	print(cp)
}

func counter(a [][]string, y, x, h, w int) string {
	dx := []int{1, 1, 1, 0, -1, -1, -1, 0}
	dy := []int{1, 0, -1, -1, -1, 0, 1, 1}
	res := 0
	for i := 0; i < 8; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if 0 <= nx && nx < w && 0 <= ny && ny < h {
			if a[ny][nx] == "#" {
				res++
			}
		}
	}
	return strconv.Itoa(res)
}

func print(a [][]string) {
	for i, row := range a {
		for j, _ := range row {
			fmt.Print(a[i][j])
		}
		fmt.Println()
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
