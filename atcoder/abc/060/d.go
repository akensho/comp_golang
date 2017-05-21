package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ns   = bufio.NewScanner(os.Stdin)
	N, W int
	w, v []int
)

func read() {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	N, _ = strconv.Atoi(row[0])
	W, _ = strconv.Atoi(row[1])
	for i := 0; i < N; i++ {
		ns.Scan()
		row = strings.Split(ns.Text(), " ")
		ww, _ := strconv.Atoi(row[0])
		vv, _ := strconv.Atoi(row[1])
		w = append(w, ww)
		v = append(v, vv)
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printTable(table [][]int) {
	for i, _ := range table {
		fmt.Println(table[i])
	}
}

func containsKey(m map[int]int, key int) bool {
	if _, ok := m[key]; ok {
		return true
	} else {
		return false
	}
}

func dfs(i, j int, dp []map[int]int) int {
	if m := dp[i]; containsKey(m, j) {
		return dp[i][j]
	}

	res := 0
	if i == 0 {
		res = 0
	}
	dp[i][j] = res
	return res
}

func main() {
	read()
	dp := make([]map[int]int, N)
	fmt.Println(N, W, w, v, dp)
	fmt.Println(dfs(0, 0, dp))
}
