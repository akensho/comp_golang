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
	INF = (1 << 32) - 1
)

func main() {
	n := intv()
	arr := make([]bool, 101)
	for i, _ := range arr {
		arr[i] = false
	}
	for i, _ := range arr {
		if i%4 == 0 {
			arr[i] = true
			for j := i + 7; j <= 100; j += 7 {
				arr[j] = true
			}
		}
		if i%7 == 0 {
			arr[i] = true
			for j := i + 4; j <= 100; j += 4 {
				arr[j] = true
			}
		}
	}
	if arr[n] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	/*
		ret1 := seven(four(n))
		ret2 := four(seven(n))
		if ret1 == 0 || ret2 == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	*/
}

func four(x int) int {
	res := x % 4
	return res
}

func seven(x int) int {
	res := x % 7
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
