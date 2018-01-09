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
	r := intSlice()
	n := r[0]
	k := r[1]
	a := intSlice()
	sort.Ints(a)
	cnt := 1
	b := make([]int, n+1)
	for i, _ := range a {
		b[a[i]]++
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] != a[i+1] {
			cnt++
		}
	}
	if cnt <= k {
		fmt.Println(0)
		return
	}
	ball := make([]Ball, 0)
	for i, v := range b {
		tmp := Ball{value: i, number: v}
		if v == 0 {
			continue
		}
		ball = append(ball, tmp)
	}
	sort.Sort(Balls(ball))
	ans := 0
	for i := 0; i < len(ball)-k; i++ {
		ans += ball[i].number
	}
	fmt.Println(ans)
}

type Ball struct {
	value  int
	number int
}

type Balls []Ball

func (b Balls) Len() int {
	return len(b)
}

func (b Balls) Less(i, j int) bool {
	return b[i].number < b[j].number
}

func (b Balls) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
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
