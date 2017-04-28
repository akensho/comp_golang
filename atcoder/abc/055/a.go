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
	x := n * 800
	y := 200 * (n / 15)
	fmt.Println(x - y)
}
