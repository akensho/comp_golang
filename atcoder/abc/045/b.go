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
	a := strings.Split(strv(), "")
	b := strings.Split(strv(), "")
	c := strings.Split(strv(), "")
	next, a := discard(a)
	winner := ""
	for {
		if next == "a" {
			if check(a) {
				winner = "A"
				break
			}
			next, a = discard(a)
		} else if next == "b" {
			if check(b) {
				winner = "B"
				break
			}
			next, b = discard(b)
		} else if next == "c" {
			if check(c) {
				winner = "C"
				break
			}
			next, c = discard(c)
		}
	}
	fmt.Println(winner)
}

func discard(p []string) (string, []string) {
	next := p[0]
	card := p[1:len(p)]
	return next, card
}

func check(p []string) bool {
	if len(p) == 0 {
		return true
	}
	return false
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
