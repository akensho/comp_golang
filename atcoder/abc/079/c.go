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
	n := strValue()
	row := strings.Split(n, "")

	a, _ := strconv.Atoi(row[0])
	b, _ := strconv.Atoi(row[1])
	c, _ := strconv.Atoi(row[2])
	d, _ := strconv.Atoi(row[3])

	if a+b+c+d == 7 {
		fmt.Printf("%d+%d+%d+%d=7\n", a, b, c, d)
	} else if a+b+c-d == 7 {
		fmt.Printf("%d+%d+%d-%d=7\n", a, b, c, d)

	} else if a+b-c+d == 7 {
		fmt.Printf("%d+%d-%d+%d=7\n", a, b, c, d)

	} else if a+b-c-d == 7 {
		fmt.Printf("%d+%d-%d-%d=7\n", a, b, c, d)

	} else if a-b+c+d == 7 {
		fmt.Printf("%d-%d+%d+%d=7\n", a, b, c, d)

	} else if a-b+c-d == 7 {
		fmt.Printf("%d-%d+%d-%d=7\n", a, b, c, d)

	} else if a-b-c+d == 7 {
		fmt.Printf("%d-%d-%d+%d=7\n", a, b, c, d)

	} else if a-b-c-d == 7 {
		fmt.Printf("%d-%d-%d-%d=7\n", a, b, c, d)
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
