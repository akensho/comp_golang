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
	queue := make([]string, 0)
	queue = append(queue, "")
	all := make([]string, 0)
	for {
		now := queue[0]
		if len(now) > 9 {
			break
		}
		queue = queue[1:]
		tail := make([]string, 3)
		tail[0] = now + "3"
		tail[1] = now + "5"
		tail[2] = now + "7"
		queue = append(queue, tail...)
		all = append(all, tail...)
	}

	ans := 0
	for _, v := range all {
		if check(v) {
			if x, _ := strconv.Atoi(v); x <= n {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func check(s string) bool {
	f1, f2, f3 := false, false, false
	for _, str := range s {
		switch str {
		case '3':
			f1 = true
		case '5':
			f2 = true
		case '7':
			f3 = true
		}
	}
	if f1 && f2 && f3 {
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

func floatv() float64 {
	return floats()[0]
}

func floats() []float64 {
	line := strs()
	slice := make([]float64, 0)
	for _, tmp := range line {
		val, err := strconv.ParseFloat(tmp, 64)
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
