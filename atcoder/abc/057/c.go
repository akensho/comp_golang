package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	ns = bufio.NewScanner(os.Stdin)
)

func read() int {
	ns.Scan()
	n, _ := strconv.Atoi(ns.Text())
	return n
}

func contains(in map[int]struct{}, key int) bool {
	if _, ok := in[key]; ok {
		return true
	} else {
		return false
	}
}

func divisors(n int) []int {
	sqrt := math.Sqrt(float64(n))
	m := make(map[int]struct{})
	for i := 1; i <= int(sqrt); i++ {
		if n%i == 0 {
			if contains(m, i) {
				continue
			} else {
				m[i] = struct{}{}
			}
			if contains(m, n/i) {
				continue
			} else {
				m[n/i] = struct{}{}
			}
		}
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func myMax(a, b int) int {
	if a < b {
		return b
	} else if a > b {
		return a
	} else {
		return a
	}
}

func myMin(a, b int) int {
	if a < b {
		return a
	} else if a > b {
		return b
	} else {
		return a
	}
}

func f(a, b int) int {
	return myMax(len(strconv.Itoa(a)), len(strconv.Itoa(b)))
}

func main() {
	n := read()
	ds := divisors(n)
	res := 2 << 29
	for _, a := range ds {
		b := n / a
		res = myMin(res, f(a, b))
	}
	fmt.Println(res)
}
