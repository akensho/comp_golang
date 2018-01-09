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

func read() (a, b int) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	a, _ = strconv.Atoi(row[0])
	b, _ = strconv.Atoi(row[1])
	return a, b
}

func main() {
	a, b := read()
	if a+b >= 10 {
		fmt.Println("error")
	} else {
		fmt.Println(a + b)
	}
}
