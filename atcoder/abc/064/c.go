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

func read() (int, []int) {
	l := intSlice()
	return l[0], intSlice()
}

func main() {
	n, a := read()
	m := map[string]int{
		"gray":   0,
		"brown":  0,
		"green":  0,
		"light":  0,
		"blue":   0,
		"yellow": 0,
		"orange": 0,
		"red":    0,
		"other":  0,
	}
	for _, rating := range a {
		if (1 <= rating) && (rating <= 399) {
			m["gray"] += 1
		} else if (400 <= rating) && (rating <= 799) {
			m["brown"] += 1
		} else if (800 <= rating) && (rating <= 1199) {
			m["green"] += 1
		} else if (1200 <= rating) && (rating <= 1599) {
			m["light"] += 1
		} else if (1600 <= rating) && (rating <= 1999) {
			m["blue"] += 1
		} else if (2000 <= rating) && (rating <= 2399) {
			m["yellow"] += 1
		} else if (2400 <= rating) && (rating <= 2799) {
			m["orange"] += 1
		} else if (2800 <= rating) && (rating <= 3199) {
			m["red"] += 1
		} else if 3200 <= rating {
			m["other"] += 1
		}
	}

	if n == m["other"] {
		fmt.Printf("%d %d\n", 1, m["other"])
		return
	}

	var num int
	for k, v := range m {
		if k == "other" {
			continue
		}
		if v > 0 {
			num++
		}
	}
	fmt.Printf("%d %d\n", num, num+m["other"])
}
