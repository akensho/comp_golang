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
	t := strv()
	if s == t && !contains(s) && !contains(t) {
		win()
		return
	}
	if s != t && !contains(s) && !contains(t) {
		lose()
		return
	}
	for i, _ := range s {
		si := s[i]
		ti := t[i]
		if si != ti {
			if si != '@' && ti != '@' {
				lose()
				return
			} else if si == '@' {
				if !check(ti) {
					lose()
					return
				}
			} else {
				if !check(si) {
					lose()
					return
				}
			}
		}
	}
	win()
}

func win() {
	fmt.Println("You can win")
}

func lose() {
	fmt.Println("You will lose")
}

func contains(x string) bool {
	for i, _ := range x {
		if x[i] == '@' {
			return true
		}
	}
	return false
}

func check(b byte) bool {
	if b == 'a' || b == 't' || b == 'c' || b == 'o' || b == 'd' || b == 'e' || b == 'r' {
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
		fmt.Printf("%#v¥n", v)
	}
}
