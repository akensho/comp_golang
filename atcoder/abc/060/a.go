package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ns = bufio.NewScanner(os.Stdin)

func read() (a, b, c string) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	return row[0], row[1], row[2]
}

func main() {
	a, b, c := read()
	if a[len(a)-1] == b[0] && b[len(b)-1] == c[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
