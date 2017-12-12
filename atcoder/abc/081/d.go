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
	_ = intValue()
	a := intSlice()
	if isSorted(a) {
		fmt.Println(0)
		return
	}
	max := max(a)
	min := min(a)
	if IntAbs(max) >= IntAbs(min) {
		p := pos(a, max)
		m, ans := sol(p, a, true)
		print(m, ans)
	} else {
		p := pos(a, min)
		m, ans := sol(p, a, false)
		print(m, ans)
	}
}

type Ans struct {
	x int
	y int
}

func print(m int, ans []Ans) {
	fmt.Println(m)
	for _, v := range ans {
		fmt.Printf("%d %d\n", v.x+1, v.y+1) // 1-indexed
	}
}

func sol(p int, a []int, b bool) (m int, ans []Ans) {
	for i, _ := range a {
		a[i] += a[p]
		m++
		ans = append(ans, Ans{x: p, y: i})
		if isSorted(a) {
			return m, ans
		}
	}
	if b {
		for i, _ := range a {
			if i == len(a)-1 {
				break
			}
			a[i+1] += a[i]
			m++
			ans = append(ans, Ans{x: i, y: i + 1})
			if isSorted(a) {
				return m, ans
			}
		}
	} else {
		for i := len(a) - 1; i >= 0; i-- {
			if i == 0 {
				break
			}
			a[i-1] += a[i]
			m++
			ans = append(ans, Ans{x: i, y: i - 1})
			if isSorted(a) {
				return m, ans
			}
		}
	}
	return m, ans
}

func pos(a []int, b int) int {
	for i, _ := range a {
		if a[i] == b {
			return i
		}
	}
	return -1
}

func isSorted(a []int) bool {
	for i, _ := range a {
		if i == len(a)-1 {
			break
		}
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func max(a []int) int {
	max := -1 << 29
	for _, v := range a {
		max = IntMax(max, v)
	}
	return max
}

func min(a []int) int {
	min := 1 << 29
	for _, v := range a {
		min = IntMin(min, v)
	}
	return min
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
