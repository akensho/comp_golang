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
)

func read() (res []uint64, k uint64) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	n, _ := strconv.Atoi(row[0])
	m, _ := strconv.Atoi(row[1])
	k = uint64(m)
	res = make([]uint64, 100000+1)
	for i := 0; i < n; i++ {
		ns.Scan()
		row = strings.Split(ns.Text(), " ")
		a, _ := strconv.Atoi(row[0])
		b, _ := strconv.Atoi(row[1])
		res[a] += uint64(b)
	}
	return res, k
}
func main() {
	ary, k := read()
	for i, v := range ary {
		if k <= v {
			fmt.Println(i)
			break
		}
		k -= v
	}
}
