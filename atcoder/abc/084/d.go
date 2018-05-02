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
	q := intv()
	l := make([]int, q+1)
	r := make([]int, q+1)
	for i := 0; i < q; i++ {
		row := ints()
		l[i], r[i] = row[0], row[1]
	}
	isPrime := eratosthenes(10000000)
	isSimillar := similar(isPrime)
	arr := make([]int, 100001)
	a := 3
	for i := 1; i < len(arr); i++ {
		if isSimillar[i] {
			arr[i] = arr[a] + 1
			a = i
		} else {
			arr[i] = arr[i-1]
		}
	}
	for i := 0; i < q; i++ {
		if isSimillar[l[i]] {
			if arr[l[i]] == 0 {
				fmt.Println(arr[r[i]])
			} else {
				fmt.Println(arr[r[i]] - arr[l[i]] + 1)
			}
		} else {
			fmt.Println(arr[r[i]] - arr[l[i]])
		}
	}
	/*
		for i := 2000; i < 2030; i++ {
			fmt.Printf("%d %d\n", i, arr[i])
		}
	*/
}

func similar(isPrime []bool) []bool {
	isSimillar := make([]bool, 100001)
	isSimillar[2017] = true
	for i, _ := range isSimillar {
		if isPrime[i] {
			if isPrime[(i+1)/2] {
				isSimillar[i] = true
			}
		}
	}
	return isSimillar
}

func eratosthenes(n int) []bool {
	isPrime := make([]bool, n+1)
	for i, _ := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := 2; i*j < len(isPrime); j++ {
				isPrime[i*j] = false
			}
		}
	}
	return isPrime
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
