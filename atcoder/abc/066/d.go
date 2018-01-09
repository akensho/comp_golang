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
	n := intSlice()[0]
	a := intSlice()
	twice := -1
	exist := make([]bool, n+2)
	for i, _ := range a {
		if exist[a[i]] {
			twice = a[i]
		} else {
			exist[a[i]] = true
		}
	}
	l := -1
	r := -1
	f := false
	for i, _ := range a {
		if f {
			if a[i] == twice {
				r = i
				break
			}
		} else {
			if a[i] == twice {
				l = i
				f = true
			}
		}
	}
	l++
	r++
	NewCombination(uint64(n + 1))
	for k := uint64(1); k <= uint64(n+1); k++ {
		all := nCm(uint64(n+1), uint64(k))
		sub := nCm(uint64(l-1+n-r+1), uint64(k-1))
		ans := ((all + UMOD) - sub) % UMOD
		fmt.Println(ans)
	}
}

func NewCombination(n uint64) {
	factorial = make([]uint64, n+1)
	inverse = make([]uint64, n+1)
	factorial[0] = 1
	inverse[0] = 1
	for i := 1; i < len(factorial); i++ {
		factorial[i] = (factorial[i-1] * uint64(i)) % UMOD
		inverse[i] = pow(factorial[i], uint64(UMOD-2)) % UMOD
	}
}

// nCm
func nCm(n, m uint64) uint64 {
	if n < m {
		return 0
	}
	return factorial[n] * inverse[m] % UMOD * inverse[n-m] % UMOD
}

// x^y
func pow(x, y uint64) (res uint64) {
	res = 1
	for y > 0 {
		if (y & 1) == 1 {
			res = (res * x) % UMOD
		}
		x = (x * x) % UMOD
		y = y >> 1
	}
	return res
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
