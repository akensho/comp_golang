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
	w, h, n := row[0], row[1], row[2]
	f := make([][]bool, h)
	for i, _ := range f {
		f[i] = make([]bool, w)
	}
	type my struct {
		x int
		y int
		a int
	}
	points := make([]my, n)
	for i, _ := range points {
		row = ints()
		points[i].x, points[i].y, points[i].a = row[0], row[1], row[2]
	}
	for _, p := range points {
		switch p.a {
		case 1:
			left(&f, p.x, h, w)
		case 2:
			f = right(f, p.x)
		case 3:
			f = down(f, p.y)
		case 4:
			f = up(f, p.y)
		}
	}
	ans := func() (res int) {
		for i, row := range f {
			for j, _ := range row {
				if f[i][j] {
					res++
				}
			}
		}
		return res
	}()
	fmt.Println(ans)
}

func left(f *[][]bool, x, h, w int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j < x {
				f[i][j] = true
			}
		}
	}
	for i, row := range f {
		for j, _ := range row {
			if j < x {
				f[i][j] = true
			}
		}
	}
}

func right(f [][]bool, x int) [][]bool {
	for i, row := range f {
		for j, _ := range row {
			if j > x {
				f[i][j] = true
			}
		}
	}
	return f
}

func up(f [][]bool, y int) [][]bool {
	for i, row := range f {
		if i > y {
			for j, _ := range row {
				f[i][j] = true
			}
		}
	}
	return f
}

func down(f [][]bool, y int) [][]bool {
	for i, row := range f {
		if i < y {
			for j, _ := range row {
				f[i][j] = true
			}
		}
	}
	return f
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
