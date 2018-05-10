package main

import (
	"bufio"
	"errors"
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
	a := intSlice()
	sort.Ints(a)
	ans := 1
	for i := 0; i < n; i++ {
		idx, err := lower_bound(a, a[i]+3)
		if err != nil {
			ans = IntMax(ans, n-i)
			continue
		}
		ans = IntMax(ans, idx-i)
	}
	fmt.Println(ans)
}

// lower_bound returns the first index of array "a".
// The index satisfied with a condition x <= a[index].
// The complexity is O(logN).
//
// Example:
//      a := []int{1, 11, 22, 33, 44, 55, 66}
//	idx, err := lower_bound(a, 36)
// Output:
//      idx is equals to 4
//
// If all elements of array "a" are smaller than x, lower_bound returns -1 and error.
//
// Example:
//      a := []int{1, 11, 22, 33, 44, 55, 66}
//	idx, err := lower_bound(a, 100)
// Output:
//      idx is equals to -1, and error is not nil

func lower_bound(a []int, x int) (int, error) {
	l, r := 0, len(a)-1
	if a[r] < x {
		e := errors.New("nothing")
		return -1, e
	}
	for l < r {
		mid := int((l + r) / 2)
		if x <= a[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l, nil
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
