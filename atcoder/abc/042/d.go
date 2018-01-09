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
	uint64MOD = uint64(1e9 + 7)
	factorial []uint64
	inverse   []uint64
)

func read() (int, int, int, int) {
	l := intSlice()
	h := l[0]
	w := l[1]
	a := l[2]
	b := l[3]
	return h, w, a, b
}

func main() {
	//h, w, a, b := read()
	fmt.Println(pow(uint64(2), uint64(16)))
}

func NewCombination(n uint64) {
	if n == 0 {
		return 1
	}
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
	return factorial[n] * inverse[m] % UMOD * inverse[n-k] % UMOD
}

// x^y
func pow(x, y uint64) (res uint64) {
	res = 1
	for y > 0 {
		if (y & 1) == 1 {
			res = (res * x) % uint64MOD
		}
		x = (x * x) % uint64MOD
		y = y >> 1
	}
	return res
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
