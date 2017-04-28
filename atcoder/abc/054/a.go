package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ns = bufio.NewScanner(os.Stdin)

func read() (a, b int) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	a, _ = strconv.Atoi(row[0])
	b, _ = strconv.Atoi(row[1])
	return a, b
}

func main() {
	a, b := read()
	if a == b {
		fmt.Println("Draw")
		return
	}
	if a == 1 {
		fmt.Println("Alice")
		return
	}
	if b == 1 {
		fmt.Println("Bob")
		return
	}
	if a > b {
		fmt.Println("Alice")
		return
	}
	if b > a {
		fmt.Println("Bob")
	}
}
