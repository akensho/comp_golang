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
	row := intSlice()
	n, m := row[0], row[1]
	adj := make([][]bool, n+1)
	for i, _ := range adj {
		adj[i] = make([]bool, n+1)
	}
	for i := 0; i < m; i++ {
		row = intSlice()
		s, t := row[0], row[1]
		adj[s][t] = true
		adj[t][s] = true
	}
	s := ""
	for i := 1; i <= n; i++ {
		s += strconv.Itoa(i)
	}
	set := make([]string, 0)
	set = append(set, s)
	r := []rune(s)
	for {
		next_permutation(r)
		if string(r)[:1] != "1" {
			break
		}
		set = append(set, string(r))
	}
	res := 0
	for _, v := range set {
		idxs := strings.Split(v, "")
		f := true
		for i := 0; i < len(idxs)-1; i++ {
			s, _ := strconv.Atoi(idxs[i])
			t, _ := strconv.Atoi(idxs[i+1])
			if adj[s][t] == false && adj[t][s] == false {
				f = false
				break
			}
		}
		if f {
			res++
		}
	}
	fmt.Println(res)
}

//
// s:= []rune("14325")
// next_permutation(s)
// fmt.Println(string(s)) => 14352
//
func next_permutation(p []rune) {
	i := len(p) - 2
	for p[i] >= p[i+1] {
		i--
	}
	j := len(p) - 1
	for p[j] <= p[i] {
		j--
	}
	p[i], p[j] = p[j], p[i]
	i, j = i+1, len(p)-1
	for i < j {
		p[i], p[j] = p[j], p[i]
		i, j = i+1, j-1
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
