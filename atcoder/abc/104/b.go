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
	s := strv()
	idx := -1
	if s[0] == 'A' {
		mid := s[2 : len(s)-1]
		cnt := 0
		for i, _ := range mid {
			if mid[i] == 'C' {
				cnt++
				idx = i + 2
			}
		}
		if cnt != 1 {
			fmt.Println("WA")
			return
		}
		for i, _ := range s {
			if i == 0 || i == idx {
				continue
			}
			if s[i]-'A' <= 25 {
				fmt.Println("WA")
				return
			}
		}
	} else {
		fmt.Println("WA")
		return
	}
	fmt.Println("AC")
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
