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

func read() (x, y int) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	x, _ = strconv.Atoi(row[0])
	y, _ = strconv.Atoi(row[1])
	return x, y
}

func add(set map[int]struct{}, list []int) {
	for _, v := range list {
		set[v] = struct{}{}
	}
}

func contains(set map[int]struct{}, key int) bool {
	if _, ok := set[key]; ok {
		return true
	} else {
		return false
	}
}

func main() {
	s1 := make(map[int]struct{})
	s2 := make(map[int]struct{})
	s3 := make(map[int]struct{})
	l1 := []int{1, 3, 5, 7, 8, 10, 12}
	l2 := []int{4, 6, 9, 11}
	l3 := []int{2}
	add(s1, l1)
	add(s2, l2)
	add(s3, l3)
	x, y := read()
	if x == 2 || y == 2 {
		fmt.Println("No")
		return
	}
	if contains(s1, x) && contains(s1, y) {
		fmt.Println("Yes")
	} else if contains(s2, x) && contains(s2, y) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
