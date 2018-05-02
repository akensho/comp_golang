package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n := intValue()
	field := make([][]int, 3)
	dp := make([][]int, 3)
	for i, _ := range field {
		field[i] = make([]int, n)
	}
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 0; i < 3; i++ {
		row := intSlice()
		sort.Ints(row)
		for j, _ := range row {
			field[i][j] = row[j]
		}
	}
	for i := 1; i <= 2; i++ {
		for j := 0; j < n; j++ {
			row := field[i-1]
			idx := binary_search(row, field[i][j])
			for k := 0; k < idx; k++ {
				dp[i][j] += dp[i-1][k]
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res += dp[2][i]
	}
	fmt.Println(res)
}

func upper_bound(a []int, x int) int {
	l, r := 0, len(a)-1
	for l < r {
		mid := int((r + l) / 2)
		if a[mid] == x {
			return mid
		} else if x < a[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

func lower_bound(a []int, x int) int {
	l, r := 0, len(a)-1
	for l < r {
		mid := int((r + l) / 2)
		if a[mid] < x {
			r = mid
		} else if x < a[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

func binary_search(a []int, x int) int {
	l, r := 0, len(a)-1
	if x < a[l] {
		return 0
	}
	if x > a[r] {
		return len(a)
	}
	mid := int((r + l) / 2)
	for l < r {
		if a[mid] == x {
			return mid
		} else if x < a[mid] {
			r = mid
		} else {
			l = mid + 1
		}
		mid = int((r + l) / 2)
	}
	return mid
}

func printField(a [][]int) {
	for i, v := range a {
		for j, _ := range v {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
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

func strValue() string {
	return strSlice()[0]
}

func strSlice() []string {
	line := strings.Split(readln(), " ")
	return line
}

func intValue() int {
	return intSlice()[0]
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
