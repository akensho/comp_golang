package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in                 = bufio.NewReader(os.Stdin)
	out                = bufio.NewWriter(os.Stdout)
	INF                = (1 << 32) - 1
	factorial, inverse []uint64
)

func main() {
	row := ints()
	w, h := row[0], row[1]
	combination(uint64(w + h - 2))
	ans := nCm(uint64(w+h-2), uint64(h-1))
	fmt.Println(ans)
}

func combination(n uint64) {
	factorial = make([]uint64, n+1)
	inverse = make([]uint64, n+1)
	factorial[0] = 1
	inverse[0] = 1
	for i := 1; i < len(factorial); i++ {
		factorial[i] = (factorial[i-1] * uint64(i)) % uint64(1e9+7)
		inverse[i] = pow(factorial[i], uint64(uint64(1e9+7)-2)) % uint64(1e9+7)
	}
}

// nCm
func nCm(n, m uint64) uint64 {
	if n < m {
		return 0
	}
	return factorial[n] * inverse[m] % uint64(1e9+7) * inverse[n-m] % uint64(1e9+7)
}

// x^y
func pow(x, y uint64) (res uint64) {
	res = 1
	for y > 0 {
		if (y & 1) == 1 {
			res = (res * x) % uint64(1e9+7)
		}
		x = (x * x) % uint64(1e9+7)
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

func floatv() float64 {
	return floats()[0]
}

func floats() []float64 {
	line := strs()
	slice := make([]float64, 0)
	for _, tmp := range line {
		val, err := strconv.ParseFloat(tmp, 64)
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

func var_dump(value ...interface{}) {
	for _, v := range value {
		fmt.Printf("%#v\n", v)
	}
}
