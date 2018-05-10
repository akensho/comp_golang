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
	INF       = 1 << 29
)

func main() {
	_ = intSlice()
	row := intSlice()
	x, y := row[0], row[1]
	a := intSlice()
	b := intSlice()
	t := a[0] + x
	ans := 0
	for {

		idx := lower_bound(b, t)
		if b[idx] < t {
			break
		}
		t = b[idx] + y
		ans++
		idx = lower_bound(a, t)
		if a[idx] < t {
			break
		}
		t = a[idx] + x
	}
	fmt.Println(ans)
}

func flight(a, b []int, path string, t int) {

}

func lower_bound(a []int, x int) int {
	l, r := 0, len(a)-1
	for l < r {
		mid := int((l + r) / 2)
		if x <= a[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
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
