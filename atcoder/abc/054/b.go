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

func read() (img, tmpl [][]string) {
	s := intSlice()
	n := s[0]
	img = make([][]string, n)
	for i, _ := range img {
		img[i] = make([]string, n)
		row := strings.Split(readln(), "")
		for j, _ := range row {
			img[i][j] = row[j]
		}
	}

	m := s[1]
	tmpl = make([][]string, m)
	for i, _ := range tmpl {
		tmpl[i] = make([]string, m)
		row := strings.Split(readln(), "")
		for j, _ := range row {
			tmpl[i][j] = row[j]
		}
	}
	return img, tmpl
}

func main() {
	img, tmpl := read()
	for i, _ := range img {
		fmt.Println(img[i])
	}
	for i, _ := range tmpl {
		fmt.Println(tmpl[i])
	}
	n := len(img)
	m := len(tmpl)
	for i := 0; i+m <= n; i++ {
		for j := 0; j+m <= n; j++ {
			for k := 0; k < m; k++ {
				if tmpl[i][j+k] == img[i][j+k] {

				}
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

func strSlice() []string {
	line := strings.Split(readln(), " ")
	return line
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
