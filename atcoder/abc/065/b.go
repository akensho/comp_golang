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

func main() {
	n := intSlice()[0]
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = intSlice()[0] - 1
	}
	b := make([]bool, n)
	idx := 0
	cnt := 0
	f := true
	for {
		if idx == 1 {
			break
		}
		if b[idx] {
			f = false
			break
		}
		b[idx] = true
		idx = a[idx]
		cnt++
	}
	if f {
		fmt.Println(cnt)
	} else {
		fmt.Println(-1)
	}
}
