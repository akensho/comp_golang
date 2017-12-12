package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	ns = bufio.NewScanner(os.Stdin)
)

func read() []int {
	ns.Scan()
	n, _ := strconv.Atoi(ns.Text())
	var s []int
	for i := 0; i < n; i++ {
		ns.Scan()
		t, _ := strconv.Atoi(ns.Text())
		s = append(s, t)
	}
	return s
}

func sol(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	if sum%10 == 0 {
		var ts []int
		for _, v := range s {
			if v%10 != 0 {
				ts = append(ts, v)
			}
		}
		if len(ts) == 0 {
			return 0
		} else {
			sort.Ints(ts)
			return sum - ts[0]
		}
	} else {
		return sum
	}
}

func main() {
	s := read()
	sort.Ints(s)
	res := sol(s)
	fmt.Println(res)
}
