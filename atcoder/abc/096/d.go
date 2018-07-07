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
)

func main() {
	n := intv()
	prime := eratosthenes(100000)
	ans := make([]int, n)
	idx := 0
	for i, p := range prime {
		if ans[len(ans)-1] != 0 {
			break
		}
		if p {
			str := strconv.Itoa(i)
			end := str[len(str)-1:]
			if end == "1" {
				ans[idx] = i
				idx++
			}
		}
	}
	for i, v := range ans {
		if i == len(ans)-1 {
			fmt.Println(v)
		} else {
			fmt.Printf("%d ", v)
		}
	}
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
