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
	s := strv()
	m := map[string]bool{"n": false, "s": false, "e": false, "w": false}
	for _, v := range s {
		switch v {
		case 'N':
			m["n"] = true
		case 'S':
			m["s"] = true
		case 'E':
			m["e"] = true
		case 'W':
			m["w"] = true
		}
	}
	if m["n"] && !m["s"] {
		fmt.Println("No")
		return
	} else if m["s"] && !m["n"] {
		fmt.Println("No")
		return
	} else if m["e"] && !m["w"] {
		fmt.Println("No")
		return
	} else if m["w"] && !m["e"] {
		fmt.Println("No")
		return
	} else {
		fmt.Println("Yes")
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
