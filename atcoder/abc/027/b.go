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
		for _, v := range a {
			res += v
		}
		return res
	}()
	if sum == 0 {
		fmt.Println(0)
		return
	}
	if sum%n != 0 {
		fmt.Println(-1)
		return
	}
	f := true
	for _, v := range a {
		if a[0] != v {
			f = false
		}
	}
	if f {
		fmt.Println(0)
		return
	}
	per := sum / n
	ans := 0
	for i := 0; i < len(a)-1; i++ {
		need := check(i, i+1, per, a)
		if need {
			ans++
		}
	}
	fmt.Println(ans)
}

func check(x, y, per int, a []int) bool {
	l := per * (x + 1)
	r := per * (len(a) - (x + 1))
	sum := func() (res int) {
		for i := 0; i <= x; i++ {
			res += a[i]
		}
		return res
	}()
	if sum != l {
		return true
	}
	sum = func() (res int) {
		for i := y; i < len(a); i++ {
			res += a[i]
		}
		return res
	}()
	if sum != r {
		return true
	}
	return false
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

func var_dump(value ...interface{}) {
	for _, v := range value {
		fmt.Printf("%#v\n", v)
	}
}
