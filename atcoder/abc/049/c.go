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
	rev := reverse(s)
	ans := rec(rev)
	if ans {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func rec(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) < 5 {
		return false
	}
	switch s[:1] {
	case "m":
		if s[:5] == "maerd" {
			return rec(s[5:])
		} else {
			return false
		}
	case "r":
		if len(s) < 6 {
			return false
		}
		if s[:6] == "resare" {
			return rec(s[6:])
		} else if s[:7] == "remaerd" {
			return rec(s[7:])
		} else {
			return false
		}
	case "e":
		if s[:5] == "esare" {
			return rec(s[5:])
		} else {
			return false
		}
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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
