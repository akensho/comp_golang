package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in                   = bufio.NewReader(os.Stdin)
	out                  = bufio.NewWriter(os.Stdout)
	MOD                  = 1e9 + 7
	UMOD                 = uint64(1e9 + 7)
	factorial            []uint64
	inverse              []uint64
	h, w, sy, sx, gy, gx int
	field                [][]string
	visited              [][]bool
)

type queue struct {
	x, y, cnt int
}

func main() {
	row := ints()
	h, w = row[0], row[1]
	row = ints()
	sy, sx = row[0]-1, row[1]-1
	row = ints()
	gy, gx = row[0]-1, row[1]-1
	field = make([][]string, h)
	visited = make([][]bool, h)
	for i, _ := range field {
		field[i] = make([]string, w)
		visited[i] = make([]bool, w)
		tmp := strv()
		tarray := strings.Split(tmp, "")
		for j := 0; j < w; j++ {
			field[i][j] = tarray[j]
		}
	}
	dx := []int{1, 0, -1, 0}
	dy := []int{0, -1, 0, 1}
	q := make([]queue, 0)
	q = append(q, queue{x: sx, y: sy, cnt: 0})
	for len(q) != 0 {
		now := q[0]
		q = q[1:]
		if visited[now.y][now.x] {
			continue
		}
		visited[now.y][now.x] = true
		if field[now.y][now.x] == "#" {
			continue
		}
		if now.x == gx && now.y == gy {
			fmt.Println(now.cnt)
			return
		}
		for i := 0; i < 4; i++ {
			nx := now.x + dx[i]
			ny := now.y + dy[i]
			if 0 <= nx && nx < w && 0 <= ny && ny < h {
				q = append(q, queue{x: nx, y: ny, cnt: now.cnt + 1})
			}
		}
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
