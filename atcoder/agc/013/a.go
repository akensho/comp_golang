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
	n := intv()
	a := ints()
	if n == 1 {
		fmt.Println(1)
		return
	}
	f := false
	for i := 0; i < n-1; i++ {
		if a[i] != a[i+1] {
			f = true
			break
		}
	}
	if !f {
		fmt.Println(0)
		return
	}
	ans := 1
	state := 'n'
	for i := 0; i < n-1; i++ {
		if state == 'n' {
			if a[i+1] > a[i] {
				state = 'i'
			} else if a[i+1] == a[i] {
				state = 'n'
			} else {
				state = 'd'
			}
		}
		if state == 'i' {
			if a[i+1] < a[i] {
				state = 'n'
				ans++
			}
		}
		if state == 'd' {
			if a[i+1] > a[i] {
				state = 'n'
				ans++
			}
		}
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
