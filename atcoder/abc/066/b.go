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
	s := strSlice()[0]
	s = cut(s)

	for {
		if len(s)%2 == 0 {
			pre := s[:len(s)/2]
			post := s[len(s)/2:]
			if pre == post {
				fmt.Println(len(s))
				break
			}
		}
		s = cut(s)
	}
}

func cut(s string) string {
	return s[:len(s)-1]
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
