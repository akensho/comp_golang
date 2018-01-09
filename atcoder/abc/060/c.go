package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in = bufio.NewReader(os.Stdin)
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

func read() (T int, t []int) {
	line := intSlice()
	T = line[1]
	t = intSlice()
	return T, t
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	T, t := read()

	x := uint64(T)
	for i := 0; i < len(t)-1; i++ {
		x += uint64(minInt(T, t[i+1]-t[i]))
	}
	fmt.Println(x)

}
