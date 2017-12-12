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
	_ = intSlice()[0]
	a := intSlice()
	var x, y, sum, res int64
	res = 1 << 40
	for _, v := range a {
		sum += int64(v)
	}
	for i := 0; i < len(a)-1; i++ {
		x += int64(a[i])
		y = sum - x
		res = Min(res, Abs(x-y))
		/*
			fmt.Println(x)
			fmt.Println(y)
			fmt.Println("x-y")
			fmt.Println(x - y)
			fmt.Println("abs")
			fmt.Println(Abs(x - y))
			fmt.Println(res)
		*/
	}
	fmt.Println(res)
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

func Min(x, y int64) int64 {
	if x < y {
		return x
	} else {
		return y
	}
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
