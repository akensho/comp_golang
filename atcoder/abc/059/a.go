package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	ns = bufio.NewScanner(os.Stdin)
)

func read() (s1, s2, s3 string) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	s1 = row[0]
	s2 = row[1]
	s3 = row[2]
	return s1, s2, s3
}

func getFirst(s string) string {
	return s[:1]
}

func main() {
	s1, s2, s3 := read()
	ans := strings.ToUpper(getFirst(s1)) + strings.ToUpper(getFirst(s2)) + strings.ToUpper(getFirst(s3))
	fmt.Println(ans)
}
