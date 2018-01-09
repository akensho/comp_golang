package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

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

var (
	sur uint64 = 1000000007
)

func f(n int) uint64 {
	if n == 0 {
		return 1
	}
	memo := make([]uint64, n+1)
	memo[0] = 1
	for i := 1; i < len(memo); i++ {
		memo[i] = (memo[i-1] * uint64(i)) % sur
	}
	return memo[n]
}

func main() {
	s := intSlice()
	n := s[0]
	m := s[1]
	if n == m {
		res := (((f(n) * f(m)) % sur) * 2) % sur
		fmt.Println(res)
	} else if math.Abs(float64(n)-float64(m)) < 2 {
		res := (f(n) * f(m)) % sur
		fmt.Println(res)
	} else {
		fmt.Println(0)
	}
}
