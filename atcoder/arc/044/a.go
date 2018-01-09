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
	isPrime := eratosthenes(n)

	if n == 1 {
		fmt.Println("Not Prime")
		return
	}

	if isPrime[n] {
		fmt.Println("Prime")
	} else {
		s := strconv.Itoa(n)
		d, _ := strconv.Atoi(s[len(s)-1:])
		if (d%2 != 0) && (d != 5) {
			sum := 0
			for _, v := range s {
				sum += int(v - '0')
			}
			if sum%3 != 0 {
				fmt.Println("Prime")

			} else {
				fmt.Println("Not Prime")
			}
		} else {
			fmt.Println("Not Prime")
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
