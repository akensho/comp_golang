package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	ns = bufio.NewScanner(os.Stdin)
)

func read() (n int, a []int) {
	ns.Scan()
	n, _ = strconv.Atoi(ns.Text())
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	for _, v := range row {
		t, _ := strconv.Atoi(v)
		a = append(a, t)
	}
	return n, a
}

func main() {
	n, a := read()
	sort.Ints(a)
	var sum int
	for i := 0; i < n-1; i++ {
		diff := a[i+1] - a[i]
		sum += diff
	}
	fmt.Println(sum)
}
