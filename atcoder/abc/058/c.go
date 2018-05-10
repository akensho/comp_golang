package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in        = bufio.NewReader(os.Stdin)
	out       = bufio.NewWriter(os.Stdout)
	MOD       = 1e9 + 7
	UMOD      = uint64(1e9 + 7)
	factorial []uint64
	inverse   []uint64
)

func main() {
	n := intValue()
	s := make([]string, n)
	for i, _ := range s {
		s[i] = strValue()
	}
	m := map[string]map[string]int{}
	for _, str := range s {
		arr := []rune(str)
		for i := 0; i < len(arr); i++ {
			cp := arr
			for j := 0; j < len(arr); j++ {
				if i == j {
					continue
				}
				cp[i], cp[j] = cp[j], cp[i]
				tmp := string(cp)
				for x := 0; x < len(tmp); x++ {
					for y := x; y < len(tmp); y++ {
						sub := tmp[x:y]
						fmt.Printf("str => %s\n", str)
						fmt.Printf("tmp => %s, sub => %s\n", tmp, sub)
						if _, ok := m[str][sub]; ok {
							m[str][sub]++
						} else {
							fmt.Println(m[str][sub])
							m[str][sub] = 1
						}
					}
				}
				//fmt.Println(string(cp))
			}
		}
	}
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

func strValue() string {
	return strSlice()[0]
}

func strSlice() []string {
	line := strings.Split(readln(), " ")
	return line
}

func intValue() int {
	return intSlice()[0]
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

func IntMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func IntMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
