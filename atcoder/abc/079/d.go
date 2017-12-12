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
	row := intSlice()
	h := row[0]
	w := row[1]
	c := make([][]int, 10)
	for i, _ := range c {
		c[i] = make([]int, 10)
	}
	for i := 0; i <= 9; i++ {
		row := intSlice()
		for j := 0; j <= 9; j++ {
			c[i][j] = row[j]
		}
	}
	a := make([][]int, h)
	for i, _ := range a {
		a[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		row := intSlice()
		for j := 0; j < w; j++ {
			a[i][j] = row[j]
		}
	}
	cost := make([]int, 10)
	for i, _ := range cost {
		cost[i] = -1
	}
	for i, _ := range cost {
		if i == -1 || i == 0 || i == 1 {
			continue
		}
		cost[i] = search(i, c)
	}
	res := 0
	for i, row := range a {
		for j, _ := range row {
			if a[i][j] == -1 {
				continue
			}
			res += cost[a[i][j]]
		}
	}
	fmt.Println(cost)
	fmt.Println(res)
}

type queue struct {
	cost int
	idx  int
}

func search(x int, field [][]int) int {
	q := make([]queue, 0)
	for i := 0; i <= 9; i++ {
		if field[x][i] < field[x][1] {
			q = append(q, queue{cost: field[x][i], idx: i})
		}
	}
	for {
		if len(q) == 0 {
			break
		}
		now := q[0]
		q = q[1:]
		if now.idx == 1 {
			if field[x][1] > now.cost {
				return now.cost
			}
		}
		for i := 0; i <= 9; i++ {
			sum := field[now.idx][i] + now.cost
			if sum < field[now.idx][1] {
				q = append(q, queue{cost: sum, idx: i})
			}
		}
	}
	return 0
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
