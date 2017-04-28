package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ns = bufio.NewScanner(os.Stdin)

func read() int {
	ns.Scan()
	n, _ := strconv.Atoi(ns.Text())
	return n
}

func main() {
	n := read()
	p := 1
	for i := 1; i <= n; i++ {
		p = p * i
		p = p % 1000000007
	}
	fmt.Println(p)
}
