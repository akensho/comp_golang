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
	a := ints()
	sum := func() (res int) {
		for i, _ := range a {
			res += a[i]
		}
		return res
	}()
	if sum%n != 0 {
		fmt.Println(-1)
		return
	}
	per := sum / n
	ans := 0
	for i := 0; i < n-1; i++ {
		l, r := 0, 0
		// left
		for j, _ := range a {
			if j > i {
				break
			}
			l += a[j]
		}
		// right
		for j, _ := range a {
			if j < i {
				continue
			}
			r += a[j]
		}
		//fmt.Printf("%d %d\n", l, r)
		if per*(i+1) == l || per*(n-(i+1)) == r {
			continue
		}
		ans++
	}
	fmt.Println(ans)
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
