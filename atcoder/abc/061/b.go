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

func read() [][]int {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	n, _ := strconv.Atoi(row[0])
	m, _ := strconv.Atoi(row[1])
	matrix := make([][]int, n+1)
	for i, _ := range matrix {
		matrix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		ns.Scan()
		row = strings.Split(ns.Text(), " ")
		a, _ := strconv.Atoi(row[0])
		b, _ := strconv.Atoi(row[1])
		matrix[a][b]++
		matrix[b][a]++
	}
	return matrix
}

func main() {
	matrix := read()
	for i, row := range matrix {
		if i == 0 {
			continue
		}
		cnt := 0
		for _, val := range row {
			cnt += val
		}
		fmt.Println(cnt)
	}
}
