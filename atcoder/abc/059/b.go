package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	ns = bufio.NewScanner(os.Stdin)
)

func read() (a, b string) {
	ns.Scan()
	a = ns.Text()
	ns.Scan()
	b = ns.Text()
	return a, b
}

func stringAt(str string, i int) string {
	return str[i : i+1]
}

func main() {
	a, b := read()
	if len(a) < len(b) {
		fmt.Println("LESS")
		return
	} else if len(a) > len(b) {
		fmt.Println("GREATER")
		return
	} else {
		for i := 0; i < len(a); i++ {
			s, _ := strconv.Atoi(stringAt(a, i))
			t, _ := strconv.Atoi(stringAt(b, i))
			if s < t {
				fmt.Println("LESS")
				return
			} else if s > t {
				fmt.Println("GREATER")
				return
			} else {
				continue
			}
		}
		fmt.Println("EQUAL")
	}
}
