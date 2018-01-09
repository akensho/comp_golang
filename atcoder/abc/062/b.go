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

func read() (h, w int, field [][]string) {
	ns.Scan()
	row := strings.Split(ns.Text(), " ")
	h, _ = strconv.Atoi(row[0])
	w, _ = strconv.Atoi(row[1])
	field = make([][]string, h)
	for i, _ := range field {
		field[i] = make([]string, w)
	}
	for i := 0; i < h; i++ {
		ns.Scan()
		row = strings.Split(ns.Text(), "")
		for j := 0; j < w; j++ {
			field[i][j] = row[j]
		}
	}
	return h, w, field
}

func main() {
	h, w, field := read()
	res := make([][]string, h+2)
	for i, _ := range res {
		res[i] = make([]string, w+2)
	}
	for i, row := range res {
		for j, _ := range row {
			res[i][j] = "#"
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			res[i+1][j+1] = field[i][j]
		}
	}
	for i, _ := range res {
		fmt.Println(strings.Join(res[i], ""))
	}
}
