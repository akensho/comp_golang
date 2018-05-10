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
	t := make([]int, n+1)
	x := make([]int, n+1)
	y := make([]int, n+1)
	t[0] = 0
	x[0] = 0
	y[0] = 0
	for i := 1; i <= n; i++ {
		row := ints()
		t[i], x[i], y[i] = row[0], row[1], row[2]
	}
	for i := 0; i < n; i++ {
		dt := t[i+1] - t[i]
		dx := abs(x[i+1] - x[i])
		dy := abs(y[i+1] - y[i])
		dp := dx + dy
		//		fmt.Printf("dy = %d\n", dy)
		//		fmt.Printf("dx = %d\n", dx)
		//		fmt.Printf("dt = %d\n", dt)
		//		fmt.Printf("dp = %d\n", dp)
		if dt == dp {
			continue
		}
		if dt > dp {
			if (dt-dp)%2 == 0 {
				continue
			}
		}
		fmt.Println("No")
		return

	}
	fmt.Println("Yes")
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
