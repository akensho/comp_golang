package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	MOD = 1e9 + 7
)

func main() {
	h, w := read()
	res := IntMin(calc(h, w), calc(w, h))
	fmt.Println(res)
}

func calc(h, w int) int {
	res := 1 << 29
	for i := 1; i <= h-1; i++ {
		s1 := i * w
		j := (int)((h - i) / 2)
		s2 := j * w
		s3 := (h - i - j) * w
		sMax := IntMax(s1, IntMax(s2, s3))
		sMin := IntMin(s1, IntMin(s2, s3))
		res = IntMin(res, IntAbs(sMax-sMin))

		j = (int)(w / 2)
		s2 = (h - i) * j
		s3 = (h - i) * (w - j)
		sMax = IntMax(s1, IntMax(s2, s3))
		sMin = IntMin(s1, IntMin(s2, s3))
		res = IntMin(res, IntAbs(sMax-sMin))
	}
	return res
}

func read() (int, int) {
	s := intSlice()
	return s[0], s[1]
}

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
