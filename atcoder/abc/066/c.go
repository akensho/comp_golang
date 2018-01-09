package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	MOD = 1e9 + 7
)

func main() {
	n := intSlice()[0]
	a := intSlice()
	e := make([]int, 0)
	o := make([]int, 0)
	for i, _ := range a {
		if i%2 == 0 {
			e = append(e, a[i])
		} else {
			o = append(o, a[i])
		}
	}
	if n%2 != 0 {
		ans := append(reverse(e), o...)
		print(ans)
	} else {
		ans := append(reverse(o), e...)
		print(ans)
	}
}

func reverse(a []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(a); i++ {
		res = append(res, a[len(a)-i-1])
	}
	return res
}

func print(a []int) {
	for i, _ := range a {
		if i == len(a)-1 {
			fmt.Println(a[i])
		} else {
			fmt.Printf("%d ", a[i])
		}
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
