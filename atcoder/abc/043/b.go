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
	s := strv()
	res := ""
	for _, v := range s {
		switch v {
		case '0':
			res += "0"
		case '1':
			res += "1"
		case 'B':
			res = delete(res)
		}
	}
	fmt.Println(res)
}

func delete(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[:len(s)-1]
}

type orignal struct {
	value string
}

type orignals []orignal

func (o orignals) Len() int {
	return len(o)
}

func (o orignals) Less(i, j int) bool {
	return o[i].value < o[j].value
}

func (o orignals) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
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
