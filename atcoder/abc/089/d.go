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
	h, w, d := row[0], row[1], row[2]
	v := h * w
	a := make([][]int, h)
	p := make(map[int]int)
	for i, _ := range a {
		a[i] = make([]int, w)
		row = ints()
		for j := 0; j < w; j++ {
			a[i][j] = row[j]
			p[row[j]] = (i * w) + j
		}
	}
	q := intv()
	l := make([]int, q)
	r := make([]int, q)
	for i := 0; i < q; i++ {
		row = ints()
		l[i], r[i] = row[0], row[1]
	}
	//
	// end of input
	//

	// init graph
	g := make([][]int, v)
	for i, _ := range g {
		g[i] = make([]int, v)
	}
	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			g[i][j] = 1 << 29
		}
	}
	for i := 0; i < v; i++ {
		for j = 0; j < v; j++ {
			// Example:
			// h = 3, w = 4, v = 12, a[0][0] ~ a[2][3], g[0][0] ~ g[11][11]
			// 1 2 6 11
			// 4 5 8 32
			// 3 7 9 13
			//
			// 2 = a[0][1] = g[1]
			// the number of vertex a[0][1] is 0*4+1 = 1 => the number of vertex a[x][y] is x*w+y
			//
			// 8 = a[1][2] = g[6]
			// the number of vertex a[1][2] is 1*4+2 = 6 => the number of vertex a[x][y] is x*w+y
			//
			// From g[i] to a[?][?] is bellow ...
			// g[6] : x = 6%4 = 1 , y = 6-(4*1) = 2 => a[1][2]
			// g[k] : x = k%w,      y = k-(k%w)     => a[k%w][k-(w*x)]

			ix := j % w
			iy := i - (w * ix)
			fmt.Printf("i = %d, j = %d, ix = %d\n", i, j, ix)
			fmt.Printf("i = %d, j = %d, iy = %d\n", i, j, iy)
			fmt.Println("====")
			/*
				if a[j%w][j-(w*(j%w))]-a[i%w][i-(w*(i%w))] == d {
					g[i][j] = abs((j%w)-(i%w)) + abs((j-(w*(j%w)))-(i-(w*(i%w))))
				}
			*/
		}
	}
	fmt.Println(d)
	for k := 0; k < v; k++ {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}
	for i := 0; i < q; i++ {
		if l[i] == r[i] {
			fmt.Println(0)
			continue
		}
		ps, pg := p[l[i]], p[r[i]]
		fmt.Println(g[ps][pg])
	}
}

func print(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%d ", m[i][j])
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
