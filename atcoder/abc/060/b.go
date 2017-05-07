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

func contains(set map[string]struct{}, key int) bool {
	if _, ok := set[strconv.Itoa(key)]; ok {
		return true
	} else {
		return false
	}
}

func main() {
	a, b, c := read()
	set := make(map[string]struct{})
	var sum, res int
	for {
		sum += a
		res := sum % b
		if res == c {
			fmt.Println("YES")
			return
		}
		if ok := contains(set, res); ok {
			fmt.Println("NO")
			return
		}
		set[strconv.Itoa(res)] = struct{}{}
	}
}
