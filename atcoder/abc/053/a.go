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
	x, _ := strconv.Atoi(ns.Text())
	return x
}

func main() {
	x := read()
	if x < 1200 {
		fmt.Println("ABC")
	} else {
		fmt.Println("ARC")
	}
}
