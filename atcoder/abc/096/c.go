package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in    = bufio.NewReader(os.Stdin)
	out   = bufio.NewWriter(os.Stdout)
	field [][]string
)

func main() {
	row := ints()
	h, w := row[0], row[1]
	field = make([][]string, h)
	for i, _ := range field {
		tmp := strv()
		tarray := strings.Split(tmp, "")
		field[i] = make([]string, w)
		for j := 0; j < w; j++ {
			field[i][j] = tarray[j]
		}
	}
	dx := []int{1, 0, -1, 0}
	dy := []int{0, -1, 0, 1}
	for i, row := range field {
		for j, value := range row {
			if value == "#" {
				f := false
				for idx := 0; idx < 4; idx++ {
					nx := j + dx[idx]
					ny := i + dy[idx]
					if 0 <= nx && nx < w {
						if 0 <= ny && ny < h {
							if field[ny][nx] == "#" {
								f = true
							}
						}
					}
				}
				if !f {
					fmt.Println("No")
					return
				}
			}
		}
	}
	fmt.Println("Yes")
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
