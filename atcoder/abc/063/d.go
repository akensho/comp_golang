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

func read() (a, b int, h []int) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	n, _ := strconv.Atoi(row[0])
	a, _ = strconv.Atoi(row[1])
	b, _ = strconv.Atoi(row[2])
	for i := 0; i < n; i++ {
		ns.Scan()
		t, _ := strconv.Atoi(ns.Text())
		h = append(h, t)
	}
	return a, b, h
}

func bs(h []int) int {
	lb := -1
	ub := len(h)
	for ub-lb > 1 {
		mid := (lb + ub) / 2

	}
}

func main() {
	a, b, h := read()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(h)
	sort.Ints(h)
	index := bs(h)
}
