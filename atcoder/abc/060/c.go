package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ns = bufio.NewScanner(os.Stdin)
	in = bufio.NewReaderSize(os.Stdin, 1000000)
)

func read() (N, T int, t []int) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	N, _ = strconv.Atoi(row[0])
	T, _ = strconv.Atoi(row[1])
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := in.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	row = strings.Split(string(buf), " ")
	for _, v := range row {
		tmp, _ := strconv.Atoi(v)
		t = append(t, tmp)
	}
	return N, T, t
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	_, T, t := read()
	x := uint64(T)
	for i := 0; i < len(t)-1; i++ {
		x += uint64(minInt(T, t[i+1]-t[i]))
	}
	fmt.Println(x)
}
