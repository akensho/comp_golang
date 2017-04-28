package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ns = bufio.NewScanner(os.Stdin)

func read() (a, b, c int) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	a, _ = strconv.Atoi(row[0])
	b, _ = strconv.Atoi(row[1])
	c, _ = strconv.Atoi(row[2])
	return a, b, c
}

func main() {
	a, b, c := read()
	if b-a == c-b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
